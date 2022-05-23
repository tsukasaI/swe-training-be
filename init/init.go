package main

import (
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Comment string `gorm:"type:varchar(200) not null"`
	UserID  uint
	User    User
}

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50) not null"`
	Email    string `gorm:"type:varchar(100) not null unique"`
	Password string `gorm:"type:varchar(255) not null"`
	Posts    []Post
	Follows  []User `gorm:"many2many:user_follows"`
}

func connect() (*gorm.DB, error) {
	dsn := "docker:docker@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func main() {
	db, err := connect()
	if err != nil {
		return
	}

	db.AutoMigrate(&User{}, &Post{})
	seedUsers()
}

func seedUsers() {
	db, err := connect()
	if err != nil {
		return
	}
	user1 := User{
		Name:     "user1",
		Email:    "user1@a.com",
		Password: "password",
		Posts: []Post{
			{Comment: "user1 comment"},
		},
		Follows: []User{
			{Name: "user2",
				Email:    "user2@a.com",
				Password: "password",
				Posts: []Post{
					{Comment: "user2 comment"},
				}},
			{Name: "user3",
				Email:    "user3@a.com",
				Password: "password",
				Posts: []Post{
					{Comment: "user3 comment"},
				}},
			{Name: "user4",
				Email:    "user4@a.com",
				Password: "password",
				Posts: []Post{
					{Comment: "user4 comment"},
				}},
			{Name: "user5",
				Email:    "user5@a.com",
				Password: "password",
				Posts: []Post{
					{Comment: "user5 comment"},
				}},
			{Name: "user6",
				Email:    "user6@a.com",
				Password: "password",
				Posts: []Post{
					{Comment: "user6 comment"},
				}},
		},
	}
	db.Create(&user1)
	var users = []User{}
	for i := 7; i < 11; i++ {
		strIte := strconv.Itoa(i)
		name := "user" + strIte
		email := "user" + strIte + "@a.com"
		comment := "user" + strIte + " comment"
		users = append(
			users,
			User{Name: name, Email: email, Password: "password", Posts: []Post{{Comment: comment}}})
	}

	db.Create(&users)
}
