package domain

import (
	"github.com/shironxn/inkarya/internal/domain/job"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name        string
	AvatarURL   string
	Location    string
	Description string
	Jobs        []job.Job
}
