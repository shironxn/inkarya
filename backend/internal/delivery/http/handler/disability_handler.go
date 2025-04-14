package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/service"
	"gorm.io/gorm"
)

type DisabilityHandler interface {
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
}

type disabilityHandler struct {
	service service.DisabilityService
}

func NewDisabilityHandler(service service.DisabilityService) DisabilityHandler {
	return &disabilityHandler{
		service: service,
	}
}

func (h *disabilityHandler) GetAll(c *fiber.Ctx) error {
	disabilities, err := h.service.GetAll()
	if err != nil {
		return err
	}

	var disabilityResponses []dto.DisabilityResponse
	for _, disability := range disabilities {
		disabilityResponses = append(disabilityResponses, dto.DisabilityResponse{
			ID:          disability.ID,
			Name:        disability.Name,
			Description: disability.Description,
			CreatedAt:   disability.CreatedAt,
			UpdatedAt:   disability.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "disabilities retrieved successfully",
		Data:    disabilityResponses,
	})
}

func (h *disabilityHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid disability id")
	}

	disability, err := h.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "disability not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "disability retrieved successfully",
		Data: dto.DisabilityResponse{
			ID:          disability.ID,
			Name:        disability.Name,
			Description: disability.Description,
			CreatedAt:   disability.CreatedAt,
			UpdatedAt:   disability.UpdatedAt,
		},
	})
}
