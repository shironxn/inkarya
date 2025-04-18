package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Name         string    `gorm:"not null"`
	Email        string    `gorm:"unique" `
	AvatarURL    string
	Bio          string
	Interest     string `gorm:"not null"`
	DOB          string `gorm:"not null"`
	Phone        string `gorm:"unique"`
	Location     string `gorm:"not null"`
	Status       string
	Availability string
	ResumeURL    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	// Many-to-Many
	Skills       []Skill      `gorm:"many2many:user_skills;"`
	Disabilities []Disability `gorm:"many2many:user_disabilities;"`

	// One-to-Many
	Posts        []Post        `gorm:"constraint:OnDelete:CASCADE;"`
	PostComments []PostComment `gorm:"constraint:OnDelete:CASCADE;"`
	PostLikes    []PostLike    `gorm:"constraint:OnDelete:CASCADE;"`

	// Many-to-Many (via pivot model)
	JobApplications   []JobApplication   `gorm:"constraint:OnDelete:CASCADE;"`
	SavedJobs         []SavedJob         `gorm:"constraint:OnDelete:CASCADE;"`
	CourseEnrollments []CourseEnrollment `gorm:"constraint:OnDelete:CASCADE;"`
}
