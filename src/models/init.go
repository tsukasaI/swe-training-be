package models

import (
	"fmt"
	"forbizbe/src/database"
	"strconv"

	"gorm.io/gorm"
)

func Initialize() {
	db, err := database.ConnectDb()
	if err != nil {
		return
	}

	if err := db.AutoMigrate(&User{}, &Post{}); err != nil {
		fmt.Println("Error occured when migrating")
		return
	}
	seedUsers(db)
}

func seedUsers(db *gorm.DB) {
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

func DropTables() {
	db, err := database.ConnectDb()
	if err != nil {
		fmt.Println("failed connect db")
		return
	}
	if err = db.Migrator().DropTable("user_follows"); err != nil {
		fmt.Println("users table drop table failed")
	}
	if err = db.Migrator().DropTable(&User{}); err != nil {
		fmt.Println("users table drop table failed")
	}
	if err = db.Migrator().DropTable(&Post{}); err != nil {
		fmt.Println("users table drop table failed")
	}
}
