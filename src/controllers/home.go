package controllers

import (
	"forbizbe/src/database"
	"forbizbe/src/models"
	"forbizbe/src/resources"
	"forbizbe/src/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	if err := validators.ValidateUserId(c); err != nil {
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
