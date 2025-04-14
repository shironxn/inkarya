package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Forum struct {
	gorm.Model
	UserID     uuid.UUID
	CategoryID uint
	Title      string `gorm:"not null"`
	Content    string `gorm:"not null"`
	Category   ForumCategory
	Comments   []ForumComment
}

type ForumCategory struct {
	gorm.Model
	Name   string  `gorm:"unique"`
	Forums []Forum `gorm:"foreignKey:CategoryID"`
}

type ForumComment struct {
	gorm.Model
	UserID  uuid.UUID
	User    User
	ForumID uint
	Content string `gorm:"not null"`
}
