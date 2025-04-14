package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/service"
	"github.com/shironxn/inkarya/pkg"
	"gorm.io/gorm"
)

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetAllUsers(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetMe(c *fiber.Ctx) error
}

type userHandler struct {
	service   service.UserService
	validator pkg.ValidatorService
	jwt       pkg.JWTService
}

func NewUserHandler(service service.UserService, validator pkg.ValidatorService, jwt pkg.JWTService) UserHandler {
	return &userHandler{
		service:   service,
		validator: validator,
		jwt:       jwt,
	}
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.UserCreateRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to parse request body")
	}

	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req.ID = userID

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "validation failed",
			Errors:  err,
		})
	}

	if err := h.service.CreateUser(req); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			switch pgErr.ConstraintName {
			case "uni_users_email":
				return fiber.NewError(fiber.StatusConflict, "user with the same email already exists")
			case "uni_users_phone":
				return fiber.NewError(fiber.StatusConflict, "user with the same phone number already exists")
			}
		}

		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusCreated,
		Message: "user created successfully",
	})
}

func (h *userHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return err
	}

	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			AvatarURL:    user.AvatarURL,
			Bio:          user.Bio,
			Interest:     user.Interest,
			DOB:          user.DOB,
			Phone:        user.Phone,
			Location:     user.Location,
			Status:       user.Status,
			Availability: user.Availability,
			ResumeURL:    user.ResumeURL,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "users retrieved successfully",
		Data:    userResponses,
	})
}

func (h *userHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id format")
	}

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "user retrieved successfully",
		Data: dto.UserResponse{
			ID:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			AvatarURL:    user.AvatarURL,
			Bio:          user.Bio,
			Interest:     user.Interest,
			DOB:          user.DOB,
			Phone:        user.Phone,
			Location:     user.Location,
			Status:       user.Status,
			Availability: user.Availability,
			ResumeURL:    user.ResumeURL,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		},
	})
}

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	var req dto.UserUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to parse request body")
	}

	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req.ID = userID

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "validation failed",
			Errors:  err,
		})
	}

	if err := h.service.UpdateUser(req); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			switch pgErr.ConstraintName {
			case "uni_users_email":
				return fiber.NewError(fiber.StatusConflict, "user with the same email already exists")
			case "uni_users_phone":
				return fiber.NewError(fiber.StatusConflict, "user with the same phone number already exists")
			}
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "user updated successfully",
	})
}

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	if err := h.service.DeleteUser(userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "user deleted successfully",
	})
}

func (h *userHandler) GetMe(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	id, err := token.Claims.GetSubject()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	userID, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id format")
	}

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "user retrieved successfully",
		Data: dto.UserResponse{
			ID:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			AvatarURL:    user.AvatarURL,
			Bio:          user.Bio,
			Interest:     user.Interest,
			DOB:          user.DOB,
			Phone:        user.Phone,
			Location:     user.Location,
			Status:       user.Status,
			Availability: user.Availability,
			ResumeURL:    user.ResumeURL,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		},
	})
}
