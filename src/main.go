package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type QueryParam struct {
	// userIdはintで指定してほしい テスト仕様書の変更したほうがよい。 取り急ぎコメント
	UserId int `form:"userId" binding:"required"`
}

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	engine.GET("/home", getHome)

	engine.Run(":8080")
}

func getHome(c *gin.Context) {
	var queryParam QueryParam
	if c.ShouldBind(&queryParam) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userを指定してください。"})
		return
	}
	userId := c.Query("userId")
	db, err := connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := findUser(db, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "userが存在しません。"})
		return
	}
	posts, err := getHomeData(db, user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func findUser(db *gorm.DB, userId string) (User, error) {
	var user User
	if err := db.Preload("Follows").First(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}


// database.go
func connect() (*gorm.DB, error) {
	dsn := "docker:docker@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

// user.go
type UserFollow struct {
	UserID   uint
	FollowId uint
}

type Post struct {
	gorm.Model
	Comment string `gorm:"type:varchar(200) not null"`
	UserID  uint
	User    User
}

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50) not null"`
	Email    string `gorm:"type:varchar(100) not null unique"`
	Password string `gorm:"type:varchar(255) not null"`
	Posts    []Post
	Follows  []User `gorm:"many2many:user_follows"`
}

// home.go
func getHomeData(db *gorm.DB, user User) ([]Post, error) {
	followIds := getFollowIds(user)

	var posts []Post
	db.Where("`user_id` in ?", followIds).Or("user_id = ?", user.ID).Preload("User").Find(&posts)

	return posts, nil
}

func getFollowIds(user User) []uint {
	var followIds []uint
	for _, followUser := range(user.Follows) {
		followIds = append(followIds, followUser.ID)
	}
	return followIds
}
