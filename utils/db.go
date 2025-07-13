package utils

import (
	"fmt"

	"github.com/karchx/goQL/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *config.Config) error {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	return nil
}
