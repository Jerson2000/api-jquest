package config

import (
	"log"
	"os"
)

var (
	JWTKey []byte
)

func configJwtKey() {
	var tempKey string
	if tempKey = os.Getenv("JWT_SECRET"); tempKey == "" {
		log.Println("env jwt secret is not set")
	}
	JWTKey = []byte(tempKey)
}
