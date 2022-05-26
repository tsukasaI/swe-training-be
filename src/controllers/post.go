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
	if err := validators.ValidatePostComment(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": resources.CreateResponseBody("InvalidField", map[string]string{"message": "入力に誤りがあります。"}),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
