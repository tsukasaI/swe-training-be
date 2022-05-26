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
			"error": resources.CreateResponseBody("UndesignatedUser", map[string]string{"message": "userを指定してください。"}),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
