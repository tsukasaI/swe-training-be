package validators

import (
	"errors"
	"unicode/utf8"
)

type BodyParamComment struct {
	Comment string `json:"comment" form:"comment" binding:"required"`
}

func ValidatePostComment(payload BodyParamComment) error {
	if utf8.RuneCountInString(payload.Comment) > 200 {
		return errors.New("文字数が超過しています。")
	}
	return nil
}
