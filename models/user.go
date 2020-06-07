package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"type:varchar(100);unique_index"`
	Password   string `gorm:"type:varchar(100)"`
	ProfilePic string `gorm:"type:varchar(1000)"`

	Posts []Post `gorm:"foreignkey:AuthorId"`
}
