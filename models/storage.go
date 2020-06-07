package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Storage struct {
}

func (storage Storage) IsUsernameTaken(username string) bool {
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var count int
	db.Model(&User{}).Where("username = ?", username).Count(&count)

	if count != 0 {
		return true
	}

	return false
}

func (storage Storage) CreateUser(username string, password string, profilePic string) error {
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	user := User{Username: username, Password: password, ProfilePic: profilePic}

	db.Save(&user)

	return nil
}
