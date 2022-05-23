package main

import (
	"forbizbe/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	// DB migration
	// models.DropTables()
	// models.Initialize()

	// main function
	engine := setUpRouter()
	engine.Run(":8080")
}

func setUpRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/home", controllers.GetHome)
	return engine
}
