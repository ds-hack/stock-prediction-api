package main

import (
	// "log"
	// "github.com/gin-gonic/gin"
	// "github.com/swaggo/files"
	// "github.com/swaggo/gin-swagger"

	// "./docs"
	"./routers"
)

func main() {
	router := routers.InitRouter()
	router.Run(":8080")
}
