package dto

import (
	"time"

	"github.com/google/uuid"
)

// Post Request DTOs
type PostCreateRequest struct {
	Title    string    `json:"title" validate:"required"`
	Content  string    `json:"content" validate:"required"`
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	ImageUrl string    `json:"image_url"`
}

type PostUpdateRequest struct {
	ID       uint      `json:"id" validate:"required"`
	Title    string    `json:"title" validate:"required"`
	Content  string    `json:"content" validate:"required"`
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	ImageUrl string    `json:"image_url"`
}

type PostDeleteRequest struct {
	ID     uint      `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

// Post Comment Request DTOs
type PostCommentCreateRequest struct {
	Content string    `json:"content" validate:"required"`
	PostID  uint      `json:"post_id" validate:"required"`
	UserID  uuid.UUID `json:"user_id" validate:"required"`
}

type PostCommentUpdateRequest struct {
	ID      uint      `json:"id" validate:"required"`
	Content string    `json:"content" validate:"required"`
	UserID  uuid.UUID `json:"user_id" validate:"required"`
}

type PostCommentDeleteRequest struct {
	ID     uint      `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

// Post Like Request DTOs
type PostLikeRequest struct {
	PostID uint      `json:"post_id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type PostUnlikeRequest struct {
	PostID uint      `json:"post_id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

// Post Response DTOs
type PostResponse struct {
	ID        uint                  `json:"id"`
	Title     string                `json:"title"`
	Content   string                `json:"content"`
	UserID    uuid.UUID             `json:"user_id"`
	ImageUrl  string                `json:"image_url"`
	User      UserBasicResponse     `json:"user"`
	Comments  []PostCommentResponse `json:"comments,omitempty"`
	Likes     int                   `json:"likes"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

type PostCommentResponse struct {
	ID        uint              `json:"id"`
	Content   string            `json:"content"`
	PostID    uint              `json:"post_id"`
	UserID    uuid.UUID         `json:"user_id"`
	User      UserBasicResponse `json:"user"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
