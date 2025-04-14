package dto

import (
	"time"

	"github.com/google/uuid"
)

type ForumCreateRequest struct {
	UserID     uuid.UUID `json:"user_id" validate:"required"`
	Title      string    `json:"title" validate:"required"`
	Content    string    `json:"content" validate:"required"`
	CategoryID uint      `json:"category_id" validate:"required"`
}

type ForumUpdateRequest struct {
	ID         uint      `json:"id"`
	UserID     uuid.UUID `json:"user_id" validate:"required"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CategoryID uint      `json:"category_id"`
}

type ForumDeleteRequest struct {
	ID     uint      `json:"id"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type ForumCommentCreateRequest struct {
	UserID  uuid.UUID `json:"user_id" validate:"required"`
	ForumID uint      `json:"forum_id" validate:"required"`
	Content string    `json:"content" validate:"required"`
}

type ForumCommentUpdateRequest struct {
	ID      uint      `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Content string    `json:"content"`
}

type ForumCommentDeleteRequest struct {
	ID     uint      `json:"id"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type ForumResponse struct {
	ID         uint                  `json:"id"`
	UserID     uuid.UUID             `json:"user_id"`
	Title      string                `json:"title"`
	Content    string                `json:"content"`
	CategoryID uint                  `json:"category_id"`
	Category   ForumCategoryResponse `json:"category"`
	User       UserBasicResponse     `json:"user"`
	CreatedAt  time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
}

type ForumCategoryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ForumCommentResponse struct {
	ID        uint              `json:"id"`
	UserID    uuid.UUID         `json:"user_id"`
	Content   string            `json:"content"`
	ForumID   uint              `json:"forum_id"`
	User      UserBasicResponse `json:"user"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
