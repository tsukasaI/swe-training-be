package models

import (
	"forbizbe/src/resources"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50) not null"`
	Email    string `gorm:"type:varchar(100) not null unique"`
	Password string `gorm:"type:varchar(255) not null"`
	Posts    []Post
	Follows  []User `gorm:"many2many:user_follows"`
}

func FindUser(db *gorm.DB, userId string) (User, error) {
	var user User
	if err := db.Preload("Follows").First(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (user *User) CreateUserResponse() resources.UserResponseForHome {
	userResponse := resources.UserResponseForHome{}
	userResponse.Id = user.ID
	userResponse.Name = user.Name
	userResponse.CreatedAt = user.CreatedAt.Format("2006/01/02/15/04/05")
	userResponse.UpdatedAt = user.UpdatedAt.Format("2006/01/02/15/04/05")
	return userResponse
}

func GetHomePosts(db *gorm.DB, user User) ([]Post, error) {
	followIds := getFollowIds(user)

	var posts []Post
	db.Where("`user_id` in ?", followIds).Or("user_id = ?", user.ID).Preload("User").Find(&posts)
	return posts, nil
}

func getFollowIds(user User) []uint {
	var followIds []uint
	for _, followUser := range user.Follows {
		followIds = append(followIds, followUser.ID)
	}
	return followIds
}
