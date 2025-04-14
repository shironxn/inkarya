package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobStatus string

const (
	Pending  JobStatus = "pending"
	Accepted JobStatus = "accepted"
	Rejected JobStatus = "rejected"
)

type Job struct {
	gorm.Model
	CompanyID   uint
	Company     Company `gorm:"foreignKey:CompanyID"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Location    string  `gorm:"not null"`
	Education   string
	SalaryMin   *int
	SalaryMax   *int

	// Many-to-Many
	SavedJobs    []SavedJob
	Applications []JobApplication
	Skills       []Skill      `gorm:"many2many:job_skills;"`
	Disabilities []Disability `gorm:"many2many:job_disabilities;"`
}

type JobApplication struct {
	gorm.Model
	UserID    uuid.UUID
	JobID     uint
	JobStatus JobStatus `gorm:"default:'pending'"`
}

type SavedJob struct {
	gorm.Model
	UserID uuid.UUID
	JobID  uint
}
