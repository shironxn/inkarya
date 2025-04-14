package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseStatus string

const (
	InActive  CourseStatus = "inactive"
	Active    CourseStatus = "active"
	Completed CourseStatus = "completed"
)

type Course struct {
	gorm.Model
	UserID      uuid.UUID
	CategoryID  uint
	Title       string
	Description string
	ImageURL    string
	Category    CourseCategory
	Enrollments []CourseEnrollment
	Lessons     []CourseLesson
}

type CourseCategory struct {
	gorm.Model
	Name    string   `gorm:"unique"`
	Courses []Course `gorm:"foreignKey:CategoryID"`
}

type CourseEnrollment struct {
	gorm.Model
	UserID   uuid.UUID
	CourseID uint
}

type CourseLesson struct {
	gorm.Model
	CourseID uint
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	Order    int
}

type UserLesson struct {
	gorm.Model
	UserID       uuid.UUID
	LessonID     uint
	CourseStatus CourseStatus `gorm:"default:'inactive'"`
}
