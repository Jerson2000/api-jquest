package config

import (
	"log"
	"os"
	"time"

	"github.com/jerson2000/jquest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func configDatabaseConnection() {
	var err error

	dsn := os.Getenv("DB_URL_STRING")
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Println(err)
	}

	sqlDB, err := Database.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	} else {
		log.Println("Failed to configure database connection pool:", err)
	}

	Database.AutoMigrate(
		&models.User{},
		&models.Company{},
		&models.Job{},
		&models.Recruiter{},
		&models.Candidate{},
		&models.Application{},
		&models.Experience{},
	)

}
