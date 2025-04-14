package domain

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name        string `gorm:"not null"`
	AvatarURL   string `gorm:"not null"`
	Location    string `gorm:"not null"`
	Description string `gorm:"not null"`
	Jobs        []Job
}
