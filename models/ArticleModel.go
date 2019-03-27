package models

import (
	"github.com/jinzhu/gorm"
)

// "article models"
type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(50)"`
	Content string `gorm:"type:text"`
	// ID      int64  `gorm:"id;primary_key;type:bigint(20);"`
}
