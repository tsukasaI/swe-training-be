package controllers

import (
	"forbizbe/src/database"
	"forbizbe/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryParam struct {
	// userIdはnumericで指定してほしい テスト仕様書の変更したほうがよい。 取り急ぎコメント
	UserId int `form:"userId" binding:"required,numeric"`
}

func GetHome(c *gin.Context) {
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
	user, err := models.FindUser(db, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "userが存在しません。"})
		return
	}
	posts, err := models.GetHomePosts(db, user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	data := models.FormHomeData(posts)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
