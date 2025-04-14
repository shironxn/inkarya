package domain

import "gorm.io/gorm"

type Disability struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string `gorm:"type:text"`
}
