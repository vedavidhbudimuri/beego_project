package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func RegisterModels() {
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})

}
