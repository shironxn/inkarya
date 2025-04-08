package config

import (
	"github.com/gofiber/fiber/v2"
)

func NewFiber(cfg *AppConfig) *fiber.App {
	return fiber.New(fiber.Config{
		AppName: cfg.Server.Name,
	})
}
