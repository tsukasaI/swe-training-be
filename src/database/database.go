package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	dsn := "docker:docker@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
