package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppSecret    string
	IsProduction string
	DBHost       string
	DBPort       string
	DBDatabase   string
	DBUsername   string
	DBPassword   string
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return &Config{}, fmt.Errorf("failed to load .env: %w", err)
	}
	return &Config{
		AppSecret:    os.Getenv("APP_SECRET"),
		IsProduction: os.Getenv("IS_PRODUCTION"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBDatabase:   os.Getenv("DB_DATABASE"),
		DBUsername:   os.Getenv("DB_USERNAME"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
	}, nil
}
