package routers

import (
	"github.com/gin-gonic/gin"
	// "github.com/swaggo/files"
	// "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.POST("/oauth2/token", api.GetAuth)

	return r
}
