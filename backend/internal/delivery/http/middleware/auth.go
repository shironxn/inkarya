package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CheckUserAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		rawUser := c.Locals("user")
		token, ok := rawUser.(*jwt.Token)
		if !ok || !token.Valid {
			return fiber.ErrUnauthorized
		}

		sub, err := token.Claims.GetSubject()
		if err != nil {
			return fiber.ErrUnauthorized
		}

		authUserID, err := strconv.Atoi(sub)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		paramID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid user ID")
		}

		if authUserID != paramID {
			return fiber.ErrForbidden
		}

		return c.Next()
	}
}
