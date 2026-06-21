package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort      string
	AppSecret    string
	IsProduction string
	DBHost       string
	DBPort       string
	DBDatabase   string
	DBUsername   string
	DBPassword   string
	DBSSLMode    string
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}
	return &Config{
		AppPort:      env("PORT", "8000"),
		AppSecret:    os.Getenv("APP_SECRET"),
		IsProduction: os.Getenv("IS_PRODUCTION"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       env("DB_PORT", "5432"),
		DBDatabase:   os.Getenv("DB_DATABASE"),
		DBUsername:   os.Getenv("DB_USERNAME"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBSSLMode:    env("DB_SSLMODE", "disable"),
	}, nil
}

func env(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
