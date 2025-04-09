package domain

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name        string
	AvatarURL   string
	Location    string
	Description string
	Jobs        []Job
}
