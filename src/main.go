package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id        int `gorm:"column:id"`
	Name      string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	dsn := "docker:docker@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Printf("%v\n", db)
	fmt.Printf("%v\n", err)

	var users []User
	db.Find(&users)
	fmt.Printf("%v\n", users)

	engine.Run(":8080")
}
