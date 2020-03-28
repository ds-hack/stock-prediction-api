package main

import (
	// "log"
	// "github.com/gin-gonic/gin"
	// "github.com/swaggo/files"
	// "github.com/swaggo/gin-swagger"

	// "./docs"
	"github.com/ds-hack/stock-prediction-api/routers"
)

func main() {
	router := routers.InitRouter()
	router.Run(":8080")
}
