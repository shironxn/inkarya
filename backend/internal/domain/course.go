package domain

import (
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
	Name     string `gorm:"unique"`
	Couerses []Course
}

type CourseEnrollment struct {
	gorm.Model
	UserID   uint
	CourseID uint
}

type CourseLesson struct {
	gorm.Model
	CourseID uint
	title    string
	content  string
	order    int
}

type UserLesson struct {
	gorm.Model
	UserID       uint
	LessonID     uint
	CourseStatus `gorm:"type:enum('inactive', 'active', 'completed');default:'inactive'"`
}
