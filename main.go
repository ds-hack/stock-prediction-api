package main

import (
	// "log"
	// "github.com/gin-gonic/gin"

	_ "github.com/ds-hack/stock-prediction-api/docs"
	"github.com/ds-hack/stock-prediction-api/routers"
)

// @title Stock Prediction API
// @version 1.0
// @description Demonstration api for dshack project

// @contact.name Contact us
// @contact.url https://datascientist-toolbox.com/
// @contact.email datascientist.toolbox.com@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	router := routers.InitRouter()

	router.Run(":8080")
}
