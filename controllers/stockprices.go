package controllers

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ds-hack/stock-prediction-api/models"
	"github.com/ds-hack/stock-prediction-api/services/stockpricesservice"
)

// GetRawStockPrices godoc
// @Summary 株価の生データを取得
// @Description 指定の企業について、株価の生データ(open/high/low/close/volume)を取得する
// @Tags stockprices
// @Produce  json
// @Param stockCode path string true "銘柄コード" minlength(4) maxlength(16)
// @Param startDate query string false "株価データ取得の開始日 YYYY-MM-DDの形式で指定"
// @Param endDate query string false "株価データ取得の終了日 YYYY-MM-DDの形式で指定"
// @Success 200 {object} models.RawStockPriceWrapper
// @Failure 400 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /stockprices/rawdata/{stockCode} [get]
func (c *Controller) GetRawStockPrices(ctx *gin.Context) {
	stockCode := ctx.Param("stockCode")
	startDate := ctx.DefaultQuery("startDate", "2019-01-01")
	endDate := ctx.DefaultQuery("endDate", time.Now().Format("2006-01-02"))
	_, err := strconv.Atoi(stockCode)
	// 銘柄コードは現状数字を想定している
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			&models.HTTPError{
				Status:  http.StatusBadRequest,
				Message: "パラメータが正しく指定されていません。",
			})
		return
	}

	rawStockPriceWrapper, err := stockpricesservice.GetRawPrices(
		stockCode, startDate, endDate,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			&models.HTTPError{
				Status:  http.StatusInternalServerError,
				Message: "株価生データの取得時にサーバーでエラーが発生しました。",
			})
		return
	}
	ctx.JSON(http.StatusOK, rawStockPriceWrapper)
}

// GetStockPricesMA godoc
// @Summary 移動平均株価データを取得
// @Description 指定の企業について、指定の計算方法による移動平均株価を取得する
// @Tags stockprices
// @Produce  json
// @Param stockCode path string true "銘柄コード" minlength(4) maxlength(16)
// @Param maTypes query string true "移動平均の計算方法 + 計算日数のセットをカンマ区切りで指定 ex) sma5,ema25"
// @Param startDate query string false "株価データ取得の開始日 YYYY-MM-DDの形式で指定"
// @Param endDate query string false "株価データ取得の終了日 YYYY-MM-DDの形式で指定"
// @Success 200 {object} models.RawStockPriceWrapper
// @Failure 400 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /stockprices/ma/{stockCode} [get]
func (c *Controller) GetStockPricesMA(ctx *gin.Context) {
	stockCode := ctx.Param("stockCode")
	maTypesQuery := ctx.Query("maTypes")
	maTypes := strings.Split(maTypesQuery, ",")
	startDate := ctx.DefaultQuery("startDate", "2019-01-01")
	endDate := ctx.DefaultQuery("endDate", time.Now().Format("2006-01-02"))
	_, err := strconv.Atoi(stockCode)
	// 銘柄コードは現状数字を想定している, MATypesに1つ以上の計算方法指定が存在すること
	if err != nil || len(maTypes) == 0 {
		ctx.JSON(http.StatusBadRequest,
			&models.HTTPError{
				Status:  http.StatusBadRequest,
				Message: "パラメータが正しく指定されていません。",
			})
		return
	}

	// 全てのMATypeが有効であること ^(sma|ema|wma)(5|25|75)$
	r := regexp.MustCompile(`^(sma|ema|wma)(5|25|75)$`)
	for _, s := range maTypes {
		if !r.MatchString(s) {
			ctx.JSON(http.StatusBadRequest,
				&models.HTTPError{
					Status:  http.StatusBadRequest,
					Message: "パラメータが正しく指定されていません。",
				})
			return
		}
	}

	stockPriceMAWrapper, err := stockpricesservice.GetMAPrices(
		stockCode, startDate, endDate, maTypes,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			&models.HTTPError{
				Status:  http.StatusInternalServerError,
				Message: "移動平均株価データの取得時にサーバーでエラーが発生しました。",
			})
		return
	}
	ctx.JSON(http.StatusOK, stockPriceMAWrapper)
}
