package domain

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID   uint
	Title    string
	Content  string
	ImageUrl string
}

type PostComment struct {
	gorm.Model
	UserID  uint
	PostID  uint
	Content string
}

type PostLike struct {
	gorm.Model
	UserID uint
	PostID uint
}
