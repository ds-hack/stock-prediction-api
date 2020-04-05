package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ds-hack/stock-prediction-api/models"
	"github.com/ds-hack/stock-prediction-api/services/companiesservice"
)

// ListCompanies godoc
// @Summary Dashboardアプリに登録されている企業情報一覧を取得
// @Description Dashboardアプリに登録されている企業の基本情報を取得します。
// @Tags companies
// @Produce  json
// @Success 200 {object} models.CompanyWrapper
// @Failure 500 {object} models.HTTPError
// @Router /companies [get]
func (c *Controller) ListCompanies(ctx *gin.Context) {
	companyWrapper, err := companiesservice.GetAll()
	if err != nil {
		// TODO: 500エラーはSlack等によりAPI管理者へ通知する
		// セキュリティ上のリスクに繋がるため、ユーザーに対してはエラーの詳細は表示しない。
		ctx.JSON(http.StatusInternalServerError,
			&models.HTTPError{
				Status:  http.StatusInternalServerError,
				Message: "企業情報一覧の取得時にサーバーでエラーが発生しました。",
			})
	}

	ctx.JSON(http.StatusOK, companyWrapper)
}

// GetCompany godoc
// @Summary 銘柄コードで指定した特定の企業情報を取得
// @Description 銘柄コードで指定した特定の企業についての詳細情報を取得する
// @Tags companies
// @Produce  json
// @Param stockCode path string true "銘柄コード" minlength(4) maxlength(16)
// @Success 200 {object} models.CompanyWrapper
// @Failure 400 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /companies [get]
func (c *Controller) GetCompany(ctx *gin.Context) {
	stockCode := ctx.Param("stockCode")
	_, err := strconv.Atoi(stockCode)
	// 銘柄コードは現状数字を想定している
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			&models.HTTPError{
				Status:  http.StatusBadRequest,
				Message: "パラメータが正しく指定されていません。",
			})
	}

	companyWrapper, err := companiesservice.GetCompanyByStockCode(stockCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			&models.HTTPError{
				Status:  http.StatusInternalServerError,
				Message: "企業情報の取得時にサーバーでエラーが発生しました。",
			})
	}

	ctx.JSON(http.StatusOK, companyWrapper)
}
