package config

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
)

// Config holds app configurations details.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	err := gotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("failed to laod env variables: %w", err)
	}

	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}

	return config, nil
}
