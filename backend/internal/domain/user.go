package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        *string `gorm:"unique"`
	AvatarURL    *string
	Bio          *string
	Interest     string
	DOB          string
	Phone        *string `gorm:"unique"`
	Location     string
	Status       *string
	Availability *string
	ResumeURL    *string

	// Many-to-Many
	Skills       []Skill      `gorm:"many2many:user_skills;"`
	Disabilities []Disability `gorm:"many2many:user_disabilities;"`

	// One-to-Many
	Posts        []Post
	PostComments []PostComment
	PostLikes    []PostLike

	// Many-to-Many (via pivot model)
	JobApplications []JobApplication
	SavedJobs       []SavedJob
}
