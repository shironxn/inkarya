package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/service"
	"github.com/shironxn/inkarya/pkg"
	"gorm.io/gorm"
)

// Forum Handler Interface
type ForumHandler interface {
	// Forum
	CreateForum(c *fiber.Ctx) error
	GetAllForums(c *fiber.Ctx) error
	GetForumByID(c *fiber.Ctx) error
	UpdateForum(c *fiber.Ctx) error
	DeleteForum(c *fiber.Ctx) error

	// Forum Category
	GetAllCategories(c *fiber.Ctx) error

	// Forum Comment
	CreateComment(c *fiber.Ctx) error
	GetCommentsByForumID(c *fiber.Ctx) error
	UpdateComment(c *fiber.Ctx) error
	DeleteComment(c *fiber.Ctx) error
}

type forumHandler struct {
	service   service.ForumService
	validator pkg.ValidatorService
	jwt       pkg.JWTService
}

func NewForumHandler(service service.ForumService, validator pkg.ValidatorService, jwt pkg.JWTService) ForumHandler {
	return &forumHandler{
		service:   service,
		validator: validator,
		jwt:       jwt,
	}
}

func (h *forumHandler) CreateForum(c *fiber.Ctx) error {
	var req dto.ForumCreateRequest

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req.UserID = userID

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to parse request body")
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "validation failed",
			Errors:  err,
		})
	}

	if err := h.service.CreateForum(req); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid category id")
		}
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusCreated,
		Message: "forum created successfully",
	})
}

func (h *forumHandler) GetAllForums(c *fiber.Ctx) error {
	result, err := h.service.GetAllForums()
	if err != nil {
		return err
	}

	var forums []dto.ForumResponse
	for _, forum := range result {
		forums = append(forums, dto.ForumResponse{
			ID:         forum.ID,
			Title:      forum.Title,
			Content:    forum.Content,
			UserID:     forum.UserID,
			CategoryID: forum.CategoryID,
			Category: dto.ForumCategoryResponse{
				ID:        forum.Category.ID,
				Name:      forum.Category.Name,
				CreatedAt: forum.Category.CreatedAt,
				UpdatedAt: forum.Category.UpdatedAt,
			},
			CreatedAt: forum.CreatedAt,
			UpdatedAt: forum.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "forums retrieved successfully",
		Data:    forums,
	})
}

func (h *forumHandler) GetForumByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid forum id")
	}

	forum, err := h.service.GetForumByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "forum not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "forum retrieved successfully",
		Data: dto.ForumResponse{
			ID:         forum.ID,
			Title:      forum.Title,
			Content:    forum.Content,
			UserID:     forum.UserID,
			CategoryID: forum.CategoryID,
			Category: dto.ForumCategoryResponse{
				ID:        forum.Category.ID,
				Name:      forum.Category.Name,
				CreatedAt: forum.Category.CreatedAt,
				UpdatedAt: forum.Category.UpdatedAt,
			},
			CreatedAt: forum.CreatedAt,
			UpdatedAt: forum.UpdatedAt,
		},
	})
}

func (h *forumHandler) UpdateForum(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid forum id")
	}

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	var req dto.ForumUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to parse request body")
	}

	req.ID = uint(id)
	req.UserID = userID

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "validation failed",
			Errors:  err,
		})
	}

	if err := h.service.UpdateForum(req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "forum not found")
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid category id")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "forum updated successfully",
	})
}

func (h *forumHandler) DeleteForum(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid forum id")
	}

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req := dto.ForumDeleteRequest{
		ID:     uint(id),
		UserID: userID,
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "validation failed",
			Errors:  err,
		})
	}

	if err := h.service.DeleteForum(req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "forum not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "forum deleted successfully",
	})
}

// Forum Category Implementation
func (h *forumHandler) GetAllCategories(c *fiber.Ctx) error {
	result, err := h.service.GetAllCategories()
	if err != nil {
		return err
	}

	var categories []dto.ForumCategoryResponse
	for _, category := range result {
		categories = append(categories, dto.ForumCategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "forum categories retrieved successfully",
		Data:    categories,
	})
}

// Forum Comment Implementation
func (h *forumHandler) CreateComment(c *fiber.Ctx) error {
	var req dto.ForumCommentCreateRequest

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}
	req.UserID = userID

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to parse request body")
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "validation failed",
			Errors:  err,
		})
	}

	if err := h.service.CreateComment(req); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid forum id or author id")
		}
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusCreated,
		Message: "forum comment created successfully",
	})
}

func (h *forumHandler) GetCommentsByForumID(c *fiber.Ctx) error {
	forumID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid forum id")
	}

	result, err := h.service.GetCommentsByForumID(uint(forumID))
	if err != nil {
		return err
	}

	var comments []dto.ForumCommentResponse
	for _, comment := range result {
		comments = append(comments, dto.ForumCommentResponse{
			ID:      comment.ID,
			Content: comment.Content,
			ForumID: comment.ForumID,
			UserID:  comment.UserID,
			User: dto.UserBasicResponse{
				ID:        comment.User.ID,
				Name:      comment.User.Name,
				AvatarURL: comment.User.AvatarURL,
			},
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "forum comments retrieved successfully",
		Data:    comments,
	})
}

func (h *forumHandler) UpdateComment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment id")
	}

	var req dto.ForumCommentUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to parse request body")
	}

	req.ID = uint(id)

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "validation failed",
			Errors:  err,
		})
	}

	if err := h.service.UpdateComment(req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "comment not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "forum comment updated successfully",
	})
}

func (h *forumHandler) DeleteComment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment id")
	}

	if err := h.service.DeleteComment(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "comment not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "forum comment deleted successfully",
	})
}
