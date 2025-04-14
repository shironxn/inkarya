package config

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
)

func NewFiber(cfg *AppConfig) *fiber.App {
	return fiber.New(fiber.Config{
		AppName: cfg.Server.Name,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if err != nil {
				return ctx.Status(code).JSON(dto.Response{
					Success: false,
					Status:  code,
					Message: err.Error(),
				})
			}

			return nil
		},
	})
}
