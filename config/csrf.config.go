package config

import (
	"log"
	"os"
)

var CSRFKey []byte

func configCSRF() {
	var tempKey string
	if tempKey = os.Getenv("CSRF_KEY"); tempKey == "" {
		log.Println("env csrf key is not set")
	}
	CSRFKey = []byte(tempKey)
}
