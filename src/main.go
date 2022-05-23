package main

import (
	"forbizbe/src/controllers"
	"forbizbe/src/database"
	"forbizbe/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryParam struct {
	// userIdはnumericで指定してほしい テスト仕様書の変更したほうがよい。 取り急ぎコメント
	UserId int `form:"userId" binding:"required,numeric"`
}

func main() {
	engine := setUpRouter()
	engine.Run(":8080")
}

func setUpRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/home", getHome)
	return engine
}

func getHome(c *gin.Context) {
	var queryParam QueryParam
	if err := c.ShouldBind(&queryParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userを指定してください。"})
		return
	}
	userId := c.Query("userId")

	db, err := database.ConnectDb()
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
	posts, err := getHomePosts(db, user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	data := formHomeData(posts)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func findUser(db *gorm.DB, userId string) (models.User, error) {
	var user models.User
	if err := db.Preload("Follows").First(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

// home.go
func getHomePosts(db *gorm.DB, user models.User) ([]models.Post, error) {
	followIds := getFollowIds(user)

	var posts []models.Post
	db.Where("`user_id` in ?", followIds).Or("user_id = ?", user.ID).Preload("User").Find(&posts)
	return posts, nil
}

func formHomeData(posts []models.Post) []controllers.PostResponse {
	var postsResponse []controllers.PostResponse
	for _, post := range posts {
		response := post.CreatePostResponse()
		postsResponse = append(postsResponse, response)
	}

	return postsResponse
}

func getFollowIds(user models.User) []uint {
	var followIds []uint
	for _, followUser := range user.Follows {
		followIds = append(followIds, followUser.ID)
	}
	return followIds
}
