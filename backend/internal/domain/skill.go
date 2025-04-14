package domain

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string `gorm:"type:text"`
}
