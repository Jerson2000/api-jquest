package config

import (
	"log"
	"os"

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
