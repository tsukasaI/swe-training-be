package models

import (
	"forbizbe/src/resources"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Comment string `gorm:"type:varchar(200) not null" json:"comment"`
	UserID  uint   `json:"-"`
	User    User   `json:"writer"`
}

func (post *Post) CreatePostResponse() resources.PostResponse {
	postResponse := resources.PostResponse{}
	postResponse.Id = post.ID
	postResponse.Comment = post.Comment
	postResponse.CreatedAt = post.CreatedAt.Format("2006/01/02/15/04/05")
	postResponse.UpdatedAt = post.UpdatedAt.Format("2006/01/02/15/04/05")
	postResponse.User = post.User.CreateUserResponse()
	return postResponse
}
