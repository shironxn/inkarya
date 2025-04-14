package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserCreateRequest struct {
	ID           uuid.UUID `json:"id" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	AvatarURL    string    `json:"avatar_url"`
	Bio          string    `json:"bio"`
	Interest     string    `json:"interest"`
	DOB          string    `json:"dob"`
	Phone        string    `json:"phone" validate:"required"`
	Location     string    `json:"location"`
	Status       string    `json:"status"`
	Availability string    `json:"availability"`
	ResumeURL    string    `json:"resume_url"`
}

type UserUpdateRequest struct {
	ID           uuid.UUID `json:"id" validate:"required"`
	Name         string    `json:"name"`
	Email        string    `json:"email" validate:"omitempty,email"`
	AvatarURL    string    `json:"avatar_url"`
	Bio          string    `json:"bio"`
	Interest     string    `json:"interest"`
	DOB          string    `json:"dob"`
	Phone        string    `json:"phone"`
	Location     string    `json:"location"`
	Status       string    `json:"status"`
	Availability string    `json:"availability"`
	ResumeURL    string    `json:"resume_url"`
}

type UserResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	AvatarURL    string    `json:"avatar_url"`
	Bio          string    `json:"bio"`
	Interest     string    `json:"interest"`
	DOB          string    `json:"dob"`
	Phone        string    `json:"phone"`
	Location     string    `json:"location"`
	Status       string    `json:"status"`
	Availability string    `json:"availability"`
	ResumeURL    string    `json:"resume_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserBasicResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatar_url"`
}
