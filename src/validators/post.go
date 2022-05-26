package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type BodyParamComment struct {
	Comment string `json:"comment" form:"comment" binding:"required"`
}

func ValidatePostComment(c *gin.Context) error {
	err := c.ShouldBindWith(&BodyParamComment{}, binding.JSON)
	if err != nil {
		return err
	}
	return nil
}
