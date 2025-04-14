package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// JWT returns a middleware that validates JWT tokens
func JWT(jwksURL string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		JWKSetURLs: []string{jwksURL},
	})
}
