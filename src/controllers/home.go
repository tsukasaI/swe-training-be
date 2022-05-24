package controllers

import (
	"forbizbe/src/database"
	"forbizbe/src/models"
	"forbizbe/src/resources"
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
		c.JSON(http.StatusNotFound, gin.H{
			"error": resources.CreateResponseBody("UndesignatedUser", map[string]string{"message": "userを指定してください。"}),
		})
		return
	}
	userId := c.Query("userId")

	db, err := database.ConnectDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": resources.CreateResponseBody("NotFoundUser", err.Error()),
		})
		return
	}
	user, err := models.FindUser(db, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": resources.CreateResponseBody("NotFoundUser", map[string]string{"message": "userが存在しません。"}),
		})
		return
	}
	posts, err := models.GetHomePosts(db, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": resources.CreateResponseBody("InternalServerError", err.Error()),
		})
		return
	}
	data := models.FormHomeData(posts)
	responseBody := resources.CreateResponseBody("ok", data)
	c.JSON(http.StatusOK, gin.H{
		"result": responseBody,
	})
}
