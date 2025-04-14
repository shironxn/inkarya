package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(cfg *AppConfig) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{})
}
