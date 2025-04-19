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
	"go.uber.org/zap"
)

type App struct {
	Fiber  *fiber.App
	Host   string
	Port   string
	Logger pkg.LoggerService
}

func NewApp() (*App, error) {
	cfg, err := config.NewAppConfig()
	if err != nil {
		return nil, err
	}

	// Initialize logger first for better error tracking
	logger := pkg.NewLogger(pkg.LoggerConfig{
		Level:       cfg.Logger.Level,
		Development: cfg.Logger.Development,
	})

	logger.Info("Initializing application",
		zap.String("app_name", cfg.Server.Name),
		zap.String("host", cfg.Server.Host),
		zap.String("port", cfg.Server.Port),
		zap.String("environment", cfg.Server.Env),
		zap.String("version", cfg.Server.Version),
	)

	db, err := config.NewGorm(cfg)
	if err != nil {
		logger.Error("Failed to initialize database", zap.Error(err))
		return nil, err
	}
	logger.Info("Database connection established")

	// Auto migrate database
	if cfg.Server.Env == "development" {
		logger.Info("Running database migrations")
		if err := db.AutoMigrate(
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
		); err != nil {
			logger.Error("Failed to run migrations", zap.Error(err))
			return nil, err
		}
		logger.Info("Database migrations completed")
	}

	app := config.NewFiber(cfg)
	validator := pkg.NewValidator()
	jwt := pkg.NewJWT()

	// Initialize repositories
	logger.Debug("Initializing repositories")
	userRepository := repository.NewUserRepository(db)
	forumRepository := repository.NewForumRepository(db)
	courseRepository := repository.NewCourseRepository(db)
	jobRepository := repository.NewJobRepository(db)
	postRepository := repository.NewPostRepository(db)
	skillRepository := repository.NewSkillRepository(db)
	disabilityRepository := repository.NewDisabilityRepository(db)

	// Initialize services
	logger.Debug("Initializing services")
	userService := service.NewUserService(userRepository, skillRepository, disabilityRepository)
	forumService := service.NewForumService(forumRepository)
	courseService := service.NewCourseService(courseRepository)
	jobService := service.NewJobService(jobRepository)
	postService := service.NewPostService(postRepository)
	skillService := service.NewSkillService(skillRepository)
	disabilityService := service.NewDisabilityService(disabilityRepository)

	// Initialize handlers
	logger.Debug("Initializing handlers")
	userHandler := handler.NewUserHandler(userService, validator, jwt)
	forumHandler := handler.NewForumHandler(forumService, validator, jwt)
	courseHandler := handler.NewCourseHandler(courseService, validator, jwt)
	jobHandler := handler.NewJobHandler(jobService, validator, jwt)
	postHandler := handler.NewPostHandler(postService, validator, jwt)
	skillHandler := handler.NewSkillHandler(skillService)
	disabilityHandler := handler.NewDisabilityHandler(disabilityService)
	healthHandler := handler.NewHealthHandler(db, cfg)

	// Setup router
	logger.Debug("Setting up router")
	router := http.NewRouter(app, cfg.Server.Version, cfg.Server.JWKSURL, &http.Handler{
		Health:     healthHandler,
		User:       userHandler,
		Forum:      forumHandler,
		Course:     courseHandler,
		Job:        jobHandler,
		Post:       postHandler,
		Skill:      skillHandler,
		Disability: disabilityHandler,
	})
	router.Setup()

	logger.Info("Application initialization completed successfully")

	return &App{
		Fiber:  app,
		Host:   cfg.Server.Host,
		Port:   cfg.Server.Port,
		Logger: logger,
	}, nil
}

func (a *App) Run() error {
	addr := a.Host + ":" + a.Port
	a.Logger.Info("Starting server",
		zap.String("address", addr),
	)
	return a.Fiber.Listen(addr)
}
