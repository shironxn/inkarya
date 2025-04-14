package dto

import (
	"time"

	"github.com/google/uuid"
)

type CourseEnrollmentCreateRequest struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	CourseID uint      `json:"course_id" validate:"required"`
}

type CourseEnrollmentDeleteRequest struct {
	ID     uint      `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type CourseResponse struct {
	ID          uint      `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CourseLessonResponse struct {
	ID        uint      `json:"id"`
	CourseID  uint      `json:"course_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CourseEnrollmentResponse struct {
	ID        uint      `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	CourseID  uint      `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
