package config

import (
	"log"
	"os"
)

var CSRFKey []byte

func configCSRF() {
	CSRFKey = []byte(os.Getenv("CSRF_KEY"))
	if len(JWTKey) == 0 {
		log.Println("env csrf key is not set", CSRFKey)
	}
}
