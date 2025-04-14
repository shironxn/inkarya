package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/domain"
	"github.com/shironxn/inkarya/internal/service"
	"github.com/shironxn/inkarya/pkg"
	"gorm.io/gorm"
)

type PostHandler interface {
	// Post
	CreatePost(c *fiber.Ctx) error
	GetAllPosts(c *fiber.Ctx) error
	GetPostByID(c *fiber.Ctx) error
	UpdatePost(c *fiber.Ctx) error
	DeletePost(c *fiber.Ctx) error

	// Post Comment
	CreateComment(c *fiber.Ctx) error
	GetCommentsByPostID(c *fiber.Ctx) error
	UpdateComment(c *fiber.Ctx) error
	DeleteComment(c *fiber.Ctx) error

	// Post Like
	LikePost(c *fiber.Ctx) error
	UnlikePost(c *fiber.Ctx) error
}

type postHandler struct {
	service   service.PostService
	validator pkg.ValidatorService
	jwt       pkg.JWTService
}

func NewPostHandler(service service.PostService, validator pkg.ValidatorService, jwt pkg.JWTService) PostHandler {
	return &postHandler{
		service:   service,
		validator: validator,
		jwt:       jwt,
	}
}

func (h *postHandler) CreatePost(c *fiber.Ctx) error {
	var req dto.PostCreateRequest

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

	if err := h.service.CreatePost(req); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid category id")
		}
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusCreated,
		Message: "post created successfully",
	})
}

func (h *postHandler) GetAllPosts(c *fiber.Ctx) error {
	result, err := h.service.GetAllPosts()
	if err != nil {
		return err
	}

	var posts []dto.PostResponse
	for _, post := range result {
		posts = append(posts, dto.PostResponse{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			UserID:   post.UserID,
			ImageUrl: post.ImageUrl,
			User: dto.UserBasicResponse{
				ID:        post.User.ID,
				Name:      post.User.Name,
				AvatarURL: post.User.AvatarURL,
			},
			Likes:     len(post.Likes),
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "posts retrieved successfully",
		Data:    posts,
	})
}

func (h *postHandler) GetPostByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post id")
	}

	post, err := h.service.GetPostByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "post not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "post retrieved successfully",
		Data: dto.PostResponse{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			UserID:   post.UserID,
			ImageUrl: post.ImageUrl,
			User: dto.UserBasicResponse{
				ID:        post.User.ID,
				Name:      post.User.Name,
				AvatarURL: post.User.AvatarURL,
			},
			Comments:  convertCommentsToResponse(post.Comments),
			Likes:     len(post.Likes),
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		},
	})
}

func (h *postHandler) UpdatePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post id")
	}

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	var req dto.PostUpdateRequest
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

	if err := h.service.UpdatePost(req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "post not found")
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
		Message: "post updated successfully",
	})
}

func (h *postHandler) DeletePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post id")
	}

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req := dto.PostDeleteRequest{
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

	if err := h.service.DeletePost(req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "post not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "post deleted successfully",
	})
}

func (h *postHandler) CreateComment(c *fiber.Ctx) error {
	var req dto.PostCommentCreateRequest

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "post not found")
		}
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusCreated,
		Message: "post comment created successfully",
	})
}

func (h *postHandler) GetCommentsByPostID(c *fiber.Ctx) error {
	postID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post id")
	}

	result, err := h.service.GetCommentsByPostID(uint(postID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "post not found")
		}
		return err
	}

	var comments []dto.PostCommentResponse
	for _, comment := range result {
		comments = append(comments, dto.PostCommentResponse{
			ID:      comment.ID,
			Content: comment.Content,
			PostID:  comment.PostID,
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
		Message: "post comments retrieved successfully",
		Data:    comments,
	})
}

func (h *postHandler) UpdateComment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment id")
	}

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	var req dto.PostCommentUpdateRequest
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

	if err := h.service.UpdateComment(req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "comment not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "post comment updated successfully",
	})
}

func (h *postHandler) DeleteComment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment id")
	}

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req := dto.PostCommentDeleteRequest{
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

	if err := h.service.DeleteComment(req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "comment not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "post comment deleted successfully",
	})
}

func (h *postHandler) LikePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post id")
	}

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req := dto.PostLikeRequest{
		UserID: userID,
		PostID: uint(id),
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "validation failed",
			Errors:  err,
		})
	}

	if err := h.service.LikePost(req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "post not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "post liked successfully",
	})
}

func (h *postHandler) UnlikePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post id")
	}

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req := dto.PostUnlikeRequest{
		UserID: userID,
		PostID: uint(id),
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "validation failed",
			Errors:  err,
		})
	}

	if err := h.service.UnlikePost(req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "post not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "post unliked successfully",
	})
}

func convertCommentsToResponse(comments []domain.PostComment) []dto.PostCommentResponse {
	var result []dto.PostCommentResponse
	for _, comment := range comments {
		result = append(result, dto.PostCommentResponse{
			ID:      comment.ID,
			Content: comment.Content,
			PostID:  comment.PostID,
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
	return result
}
