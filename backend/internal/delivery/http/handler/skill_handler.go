package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/service"
	"gorm.io/gorm"
)

type SkillHandler interface {
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
}

type skillHandler struct {
	service service.SkillService
}

func NewSkillHandler(service service.SkillService) SkillHandler {
	return &skillHandler{
		service: service,
	}
}

func (h *skillHandler) GetAll(c *fiber.Ctx) error {
	skills, err := h.service.GetAll()
	if err != nil {
		return err
	}

	var skillResponses []dto.SkillResponse
	for _, skill := range skills {
		skillResponses = append(skillResponses, dto.SkillResponse{
			ID:          skill.ID,
			Name:        skill.Name,
			Description: skill.Description,
			CreatedAt:   skill.CreatedAt,
			UpdatedAt:   skill.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "skills retrieved successfully",
		Data:    skillResponses,
	})
}

func (h *skillHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid skill id")
	}

	skill, err := h.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "skill not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "skill retrieved successfully",
		Data: dto.SkillResponse{
			ID:          skill.ID,
			Name:        skill.Name,
			Description: skill.Description,
			CreatedAt:   skill.CreatedAt,
			UpdatedAt:   skill.UpdatedAt,
		},
	})
}
