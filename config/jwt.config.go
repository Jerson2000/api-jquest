package config

import (
	"log"
	"os"
)

var (
	JWTKey []byte
)

func configJwtKey() {
	tempKey := os.Getenv("JWT_SECRET")
	if tempKey == "" {
		log.Fatal("FATAL: JWT_SECRET environment variable is not set")
	}
	JWTKey = []byte(tempKey)
}
