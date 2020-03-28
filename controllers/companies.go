package controllers

import (
	"net/http"

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
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
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
