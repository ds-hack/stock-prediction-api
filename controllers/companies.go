package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListCompanies godoc
// @Summary Dashboardアプリに登録されている企業情報一覧を取得
// @Description Dashboardアプリに登録されている企業の基本情報を取得します。
// @Tags companies
// @Produce  json
// @Success 200 {object} model.CompanyWrapper
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /companies [get]
func ListCompanies(c *gin.Context) {
	// TODO: implement logic!!
	c.JSON(http.StatusOK, gin.H{"message": "do some magic!!"})
}
