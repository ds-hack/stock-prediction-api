package routers

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/ds-hack/stock-prediction-api/controllers"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	c := controllers.NewController()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		companies := v1.Group("/companies")
		{
			companies.GET("", c.ListCompanies)
		}

		stockprices := v1.Group("/stockprices")
		{
			stockprices.GET("rawdata/:stockCode", c.GetRawStockPrices)
		}
	}

	return r
}
