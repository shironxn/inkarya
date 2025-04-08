package handler

import "github.com/gofiber/fiber/v2"

type HealthHandler interface {
	Check(c *fiber.Ctx) error
}

type healthHandler struct{}

func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}

func (h *healthHandler) Check(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}
