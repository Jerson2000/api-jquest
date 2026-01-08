package config

import (
	"log"
	"os"
)

var (
	JWTKey []byte
)

func configJwtKey() {
	JWTKey := []byte(os.Getenv("JWT_SECRET"))
	if len(JWTKey) == 0 {
		log.Println("env jwt secret is not set", JWTKey)
	}
}
