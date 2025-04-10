package http

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shironxn/inkarya/internal/delivery/http/handler"
	"github.com/shironxn/inkarya/internal/delivery/http/middleware"
)

type Router struct {
	App     *fiber.App
	Version string
	JWKSURL string
}

type Handler struct {
	Health handler.HealthHandler
	User   handler.UserHandler
}

func NewRouter(app *fiber.App, version string, jwksURL string) *Router {
	return &Router{
		App:     app,
		Version: version,
		JWKSURL: jwksURL,
	}
}

func (r *Router) Setup(handler *Handler) {
	public := r.App.Group("/api/v" + r.Version)

	public.Get("/health", handler.Health.Check)

	public.Post("/users", handler.User.Create)
	public.Get("/users", handler.User.GetAll)
	public.Get("/users/:id", handler.User.GetByID)
	public.Put("/users/:id", middleware.CheckUserAuth(), handler.User.Update)
	public.Delete("/users/:id", middleware.CheckUserAuth(), handler.User.Delete)

	private := r.App.Group("/api/v"+r.Version, jwtware.New(jwtware.Config{
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
