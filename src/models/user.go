package models

import (
	"forbizbe/src/controllers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50) not null" json:"name"`
	Email    string `gorm:"type:varchar(100) not null unique" json:"email"`
	Password string `gorm:"type:varchar(255) not null" json:"-"`
	Posts    []Post `json:"posts"`
	Follows  []User `gorm:"many2many:user_follows"`
}

func (user *User) CreateUserResponse() controllers.UserResponse {
	userResponse := controllers.UserResponse{}
	userResponse.Id = user.ID
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.CreatedAt = user.CreatedAt.Format("2006/01/02/15/04/05")
	userResponse.UpdatedAt = user.UpdatedAt.Format("2006/01/02/15/04/05")
	return userResponse
}
