package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Post struct {
	gorm.Model

	Content  string `gorm:"type:varchar(1000)"`
	PostedAt time.Time
	AuthorId int
}
