package validators

import (
	"errors"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type BodyParamComment struct {
	Comment string `json:"comment" form:"comment" binding:"required"`
}

func ValidatePostComment(c *gin.Context) error {
	var payload BodyParamComment
	err := c.ShouldBindWith(&payload, binding.JSON)
	if err != nil {
		return err
	}

	if err := validatePostComment(payload); err != nil {
		return err
	}
	return nil
}

func validatePostComment(payload BodyParamComment) error {
	if utf8.RuneCountInString(payload.Comment) > 100 {
		return errors.New("文字数が超過しています。")
	}
	return nil
}
