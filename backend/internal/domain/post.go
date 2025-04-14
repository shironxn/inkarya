package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID   uuid.UUID
	User     User
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	ImageUrl string
	Comments []PostComment
	Likes    []PostLike
}

type PostComment struct {
	gorm.Model
	UserID  uuid.UUID
	User    User
	PostID  uint
	Content string `gorm:"not null"`
}

type PostLike struct {
	gorm.Model
	UserID uuid.UUID
	PostID uint
}
