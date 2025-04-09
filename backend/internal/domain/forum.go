package domain

import "gorm.io/gorm"

type Forum struct {
	gorm.Model
	UserID     uint
	CategoryID uint
	Title      string
	Content    string
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
	UserID  uint
	ForumID uint
	Content string
}
