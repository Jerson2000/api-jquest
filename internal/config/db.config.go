package config

import (
	"log"
	"os"

	"github.com/jerson2000/jquest/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func configDatabaseConnection() {
	var err error

	dsn := os.Getenv("DB_URL_STRING")
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatal(err)
	}

	Database.AutoMigrate(
		&models.User{},
	)

}
