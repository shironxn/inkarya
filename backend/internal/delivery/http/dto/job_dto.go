package dto

import (
	"time"

	"github.com/google/uuid"
)

type JobApplicationRequest struct {
	JobID  uint      `json:"job_id"`
	UserID uuid.UUID `json:"user_id"`
}

type JobApplicationResponse struct {
	ID        uint      `json:"id"`
	JobID     uint      `json:"job_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type JobSaveRequest struct {
	JobID  uint      `json:"job_id"`
	UserID uuid.UUID `json:"user_id"`
}
