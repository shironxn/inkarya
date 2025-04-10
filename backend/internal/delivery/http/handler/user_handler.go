package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/service"
	"github.com/shironxn/inkarya/pkg"
	"gorm.io/gorm"
)

type UserHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type userHandler struct {
	service   service.UserService
	validator *pkg.Validator
}

func NewUserHandler(service service.UserService, validator *pkg.Validator) UserHandler {
	return &userHandler{
		service:   service,
		validator: validator,
	}
}

func (h *userHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "Validation failed",
			Errors:  err,
		})
	}

	if err := h.service.Register(req); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			switch pgErr.ConstraintName {
			case "uni_users_email":
				return fiber.NewError(fiber.StatusConflict, "User with the same email already exists")
			case "uni_users_phone":
				return fiber.NewError(fiber.StatusConflict, "User with the same phone number already exists")
			}
		}

		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusCreated,
		Message: "User created successfully",
	})
}

func (h *userHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.service.GetAll()
	if err != nil {
		return err
	}

	var users []dto.UserResponse
	for _, user := range result {
		users = append(users, dto.UserResponse{
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
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

func (h *userHandler) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}

		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "User retrieved successfully",
		Data: dto.UserResponse{
			ID:           data.ID,
			Name:         data.Name,
			Email:        data.Email,
			AvatarURL:    data.AvatarURL,
			Bio:          data.Bio,
			Interest:     data.Interest,
			DOB:          data.DOB,
			Phone:        data.Phone,
			Location:     data.Location,
			Status:       data.Status,
			Availability: data.Availability,
			ResumeURL:    data.ResumeURL,
		},
	})
}

func (h *userHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	req.ID = uint(id)

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "Validation failed",
			Errors:  err,
		})
	}

	if err := h.service.Update(req); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			switch pgErr.ConstraintName {
			case "uni_users_email":
				return fiber.NewError(fiber.StatusConflict, "User with the same email already exists")
			case "uni_users_phone":
				return fiber.NewError(fiber.StatusConflict, "User with the same phone number already exists")
			}
		}

		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "User updated successfully",
	})
}

func (h *userHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}

	if err := h.service.Delete(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}

		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "User deleted successfully",
	})
}
