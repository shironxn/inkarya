package domain

import "gorm.io/gorm"

type JobStatus string

const (
	Pending  JobStatus = "pending"
	Accepted JobStatus = "accepted"
	Rejected JobStatus = "rejected"
)

type Job struct {
	gorm.Model
	CompanyID   uint
	Title       string
	Description string
	Location    string
	Education   string
	SalaryMin   *int
	SalaryMax   *int

	// Many-to-Many
	SavedJobs    []SavedJob
	Applications []JobApplication
}

type JobApplication struct {
	gorm.Model
	UserID    uint
	JobID     uint
	JobStatus JobStatus `gorm:"default:'pending'"`
}

type SavedJob struct {
	gorm.Model
	UserID uint
	JobID  uint
}
