package controllers

import (
	"forbizbe/src/resources"
	"forbizbe/src/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostPost(c *gin.Context) {
	if err := validators.ValidateUserId(c); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": resources.CreateResponseBody("UndesignatedUser", "ユーザーを指定してください。")
		})
	}
	c.JSON(http.StatusOK, gin.H{})
}
