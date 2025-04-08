package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shironxn/inkarya/internal/config"
	"github.com/shironxn/inkarya/internal/delivery/http"
	"github.com/shironxn/inkarya/internal/delivery/http/handler"
)

type App struct {
	Fiber *fiber.App
	Port  string
}

func NewApp() (*App, error) {
	cfg, err := config.NewAppConfig()
	if err != nil {
		return nil, err
	}

	app := config.NewFiber(cfg)

	healthHandler := handler.NewHealthHandler()

	router := http.NewRouter(app, cfg.Server.JWKSURL)
	router.Setup(&http.Handler{
		Health: healthHandler,
	})

	return &App{
		Fiber: app,
		Port:  cfg.Server.Port,
	}, nil
}

func (a *App) Run() error {
	return a.Fiber.Listen(":" + a.Port)
}
