package main

import (
	"fmt"
	"forbizbe/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	// DB migration
	// models.DropTables()
	// models.Initialize()

	// main function
	engine := setUpRouter()
	err := engine.Run(":8080")
	if err != nil {
		fmt.Printf("Error occured when starting HTTP server: %v", err)
	}
}

func setUpRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/home", controllers.GetHome)
	return engine
}
