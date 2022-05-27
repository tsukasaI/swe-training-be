package controllers

import (
	"forbizbe/src/database"
	"forbizbe/src/models"
	"forbizbe/src/resources"
	"forbizbe/src/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostPost(c *gin.Context) {
	if err := validators.ValidateUserId(c); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": resources.CreateResponseBody("UndesignatedUser", map[string]string{"message": "userを指定してください。"}),
		})
		return
	}
	if err := validators.ValidatePostComment(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": resources.CreateResponseBody("InvalidField", map[string]string{"message": "入力に誤りがあります。"}),
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
	_, err = models.FindUser(db, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": resources.CreateResponseBody("NotFoundUser", map[string]string{"message": "userが存在しません。"}),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
