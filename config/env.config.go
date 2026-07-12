package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	AppEnv         string
	AllowedOrigins []string
)

func configLoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found — using system environment variables")
	}

	AppEnv = os.Getenv("APP_ENV")
	if AppEnv == "" {
		AppEnv = "development"
	}

	originsStr := os.Getenv("ALLOWED_ORIGINS")
	if originsStr != "" {
		for _, origin := range strings.Split(originsStr, ",") {
			trimmed := strings.TrimSpace(origin)
			if trimmed != "" {
				AllowedOrigins = append(AllowedOrigins, trimmed)
			}
		}
	}

	// Default development origins if none are specified
	if len(AllowedOrigins) == 0 {
		AllowedOrigins = []string{
			"http://localhost:3000",
			"http://localhost:5173",
			"http://localhost:8080",
		}
	}
}

