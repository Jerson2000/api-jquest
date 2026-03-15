package config

import (
	"fmt"
	"os"
	"time"

	"github.com/jerson2000/jquest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func configDatabaseConnection() error {
	dsn := os.Getenv("DB_URL_STRING")
	var err error

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	sqlDB, err := Database.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Verify connection is healthy
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	// Perform migrations only if connection is successful
	if err := Database.AutoMigrate(
		&models.User{},
		&models.Company{},
		&models.Job{},
		&models.Recruiter{},
		&models.Candidate{},
		&models.Application{},
		&models.Experience{},
	); err != nil {
		return fmt.Errorf("database migration failed: %w", err)
	}

	return nil
}
