package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	UserID    uint
	Path      string
	Likes     int
	Downloads int
	Comments  []Comment
}

type Comment struct {
	gorm.Model
	ImageID uint
	Text    string
}
