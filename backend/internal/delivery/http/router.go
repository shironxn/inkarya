package http

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shironxn/inkarya/internal/delivery/http/handler"
)

type Router struct {
	App     *fiber.App
	JWKSURL string
}

type Handler struct {
	Health handler.HealthHandler
}

func NewRouter(app *fiber.App, jwksURL string) *Router {
	return &Router{
		App:     app,
		JWKSURL: jwksURL,
	}
}

func (r *Router) Setup(handler *Handler) {
	public := r.App.Group("/")

	public.Get("/health", handler.Health.Check)

	private := r.App.Group("/api/v1", jwtware.New(jwtware.Config{
		JWKSetURLs: []string{r.JWKSURL},
	}))

	private.Get("/me", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		id, _ := user.Claims.GetSubject()
		return c.JSON(fiber.Map{
			"message": "Welcome, you're logged in!",
			"user":    user,
			"id":      id,
		})
	})
}
