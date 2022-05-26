package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type QueryParamUserId struct {
	UserId int `json:"userId" form:"userId" binding:"required,numeric"`
}

func ValidateUserId(c *gin.Context) error {
	return c.ShouldBindWith(&QueryParamUserId{}, binding.Query)
}
