package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		db, err := connect()
		fmt.Println(db)
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	engine.GET("/home", func(c *gin.Context) {
		userId := c.Query("userId")
		fmt.Printf("%v\n", userId)
		posts, err := getHome(userId)
		fmt.Printf("%v\n", err)
		fmt.Printf("%v\n", posts)
		c.JSON(http.StatusOK, gin.H{
			"data": posts,
		})
	})

	engine.Run(":8080")
}

// database.go
func connect() (*gorm.DB, error) {
	dsn := "docker:docker@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

// user.go
type UserFollow struct {
	UserID uint
	FollowId uint
}

type Post struct {
	Id uint `gorm:"primaryKey"`
	Comment string
	UserID uint
	User User
	CreatedAt string
	UpdatedAt string
}

type User struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}


// home.go
func getHome (userId string) ([]Post, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	var posts []Post
	subQuery := db.Select("`follow_id`").Where("user_id = ?", userId).Table("user_follows")
	db.Where("`user_id` in (?)", subQuery).Or("user_id = ?", userId).Preload("User").Find(&posts)
	fmt.Printf("%v\n", posts)
	return posts, nil
}
