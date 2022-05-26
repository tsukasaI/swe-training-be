package validators

import "github.com/gin-gonic/gin"

type QueryParamUserId struct {
	UserId int `form:"userId" binding:"required,numeric"`
}

func ValidateUserId(c *gin.Context) error {
	return c.ShouldBind(&QueryParamUserId{})
}
