package config

import (
	"log"
	"os"
)

var CSRFKey []byte

func configCSRF() {
	tempKey := os.Getenv("CSRF_KEY")
	if tempKey == "" {
		log.Fatal("FATAL: CSRF_KEY environment variable is not set")
	}
	CSRFKey = []byte(tempKey)
}
