package controllers

import (
	"forbizbe/src/database"
	"forbizbe/src/models"
	"forbizbe/src/resources"
	"forbizbe/src/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func PostPost(c *gin.Context) {
	if err := validators.ValidateUserId(c); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": resources.CreateResponseBody("UndesignatedUser", map[string]string{"message": "userを指定してください。"}),
		})
		return
	}

	var payload validators.BodyParamComment
	err := c.ShouldBindWith(&payload, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": resources.CreateResponseBody("InvalidField", map[string]string{"message": "入力に誤りがあります。"}),
		})
		return
	}

	if err := validators.ValidatePostComment(payload); err != nil {
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
	user, err := models.FindUser(db, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": resources.CreateResponseBody("NotFoundUser", map[string]string{"message": "userが存在しません。"}),
		})
		return
	}

	// todo postの保存, postの取り出し
	post, err := models.CreatePost(db, userId, payload.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": resources.CreateResponseBody("InertnalServerError", map[string]string{"message": err.Error()}),
		})
		return
	}

	post.User = user
	response := post.CreatePostResponse()

	c.JSON(http.StatusOK, gin.H{
		"result": resources.CreateResponseBody("ok", map[string]resources.PostResponse{"data": response}),
	})
}
