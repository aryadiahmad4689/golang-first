package models

import (
	"github.com/jinzhu/gorm"
)

// "article models"
type Article struct {
	gorm.Model
	ID      int64  `json:"id"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
