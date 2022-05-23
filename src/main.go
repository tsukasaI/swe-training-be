package main

import (
	"forbizbe/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := setUpRouter()
	engine.Run(":8080")
}

func setUpRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/home", controllers.GetHome)
	return engine
}
