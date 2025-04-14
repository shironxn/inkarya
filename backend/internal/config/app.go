package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logger   LoggerConfig
}

type ServerConfig struct {
	Name    string
	Env     string
	Host    string
	Port    string
	Version string
	JWKSURL string
}

type DatabaseConfig struct {
	DSN string
}

type LoggerConfig struct {
	Level       string
	Development bool
}

func NewAppConfig() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &AppConfig{
		Server: ServerConfig{
			Name:    os.Getenv("APP_NAME"),
			Env:     os.Getenv("APP_ENV"),
			Host:    os.Getenv("APP_HOST"),
			Port:    os.Getenv("APP_PORT"),
			Version: os.Getenv("APP_VERSION"),
			JWKSURL: os.Getenv("JWKS_URL"),
		},
		Database: DatabaseConfig{
			DSN: os.Getenv("DB_DSN"),
		},
		Logger: LoggerConfig{
			Level:       os.Getenv("LOG_LEVEL"),
			Development: os.Getenv("APP_ENV") == "development",
		},
	}, nil
}
