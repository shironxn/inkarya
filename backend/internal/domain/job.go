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
}

type JobApplications struct {
	gorm.Model
	UserID       uint
	JobID        uint
	CourseStatus `gorm:"type:enum('pending', 'accepted', 'rejected');default:'pending'"`
}
