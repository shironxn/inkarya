package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shironxn/inkarya/internal/config"
	"github.com/shironxn/inkarya/internal/delivery/http"
	"github.com/shironxn/inkarya/internal/delivery/http/handler"
	"github.com/shironxn/inkarya/internal/domain"
	"github.com/shironxn/inkarya/internal/repository"
	"github.com/shironxn/inkarya/internal/service"
	"github.com/shironxn/inkarya/pkg"
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

	db, err := config.NewGorm(cfg)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		// User
		&domain.User{},
		&domain.Skill{},
		&domain.Disability{},

		// Post
		&domain.Post{},
		&domain.PostLike{},
		&domain.PostComment{},

		// Forum
		&domain.Forum{},
		&domain.ForumComment{},
		&domain.ForumCategory{},

		// Company & Job
		&domain.Company{},
		&domain.Job{},
		&domain.SavedJob{},
		&domain.JobApplication{},

		// Course
		&domain.Course{},
		&domain.CourseCategory{},
		&domain.CourseEnrollment{},
		&domain.CourseLesson{},
		&domain.UserLesson{},
	)

	app := config.NewFiber(cfg)
	validator := pkg.NewValidator()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, validator)

	healthHandler := handler.NewHealthHandler()

	router := http.NewRouter(app, cfg.Server.Version, cfg.Server.JWKSURL)
	router.Setup(&http.Handler{
		Health: healthHandler,
		User:   userHandler,
	})

	return &App{
		Fiber: app,
		Port:  cfg.Server.Port,
	}, nil
}

func (a *App) Run() error {
	return a.Fiber.Listen(":" + a.Port)
}
