package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shironxn/inkarya/internal/config"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"gorm.io/gorm"
)

type HealthHandler interface {
	Check(c *fiber.Ctx) error
}

type healthHandler struct {
	db     *gorm.DB
	config *config.AppConfig
}

func NewHealthHandler(db *gorm.DB, cfg *config.AppConfig) HealthHandler {
	return &healthHandler{
		db:     db,
		config: cfg,
	}
}

func (h *healthHandler) Check(c *fiber.Ctx) error {
	start := time.Now()

	// Check database connection
	dbStatus := "healthy"
	sqlDB, err := h.db.DB()
	if err != nil {
		dbStatus = "unhealthy"
	} else {
		if err := sqlDB.Ping(); err != nil {
			dbStatus = "unhealthy"
		}
	}

	// Create health status
	healthStatus := dto.HealthStatus{
		Status:      "healthy",
		Environment: h.config.Server.Env,
		Version:     h.config.Server.Version,
		Timestamp:   time.Now(),
		Database:    dbStatus,
	}

	// If database is unhealthy, mark overall status as unhealthy
	if dbStatus == "unhealthy" {
		healthStatus.Status = "unhealthy"
	}

	// Calculate response time
	responseTime := time.Since(start)

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "health check successfully",
		Data: dto.HealthResponse{
			Status:       healthStatus,
			ResponseTime: responseTime.String(),
		},
	})
}
