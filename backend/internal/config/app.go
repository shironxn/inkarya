package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Name    string
	Env     string
	Port    string
	Version string
	JWKSURL string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func NewAppConfig() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &AppConfig{
		Server: ServerConfig{
			Name:    os.Getenv("APP_NAME"),
			Env:     os.Getenv("APP_ENV"),
			Port:    os.Getenv("APP_PORT"),
			Version: os.Getenv("APP_VERSION"),
			JWKSURL: os.Getenv("JWKS_URL"),
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
	}, nil
}
