package config

import (
	"log/slog"
	"os"
)

func InitConfig() {
	InitLogger()
	configLoadEnv()
	configJwtKey()
	configCSRF()

	if err := configDatabaseConnection(); err != nil {
		slog.Error("database connection failed", "error", err)
		os.Exit(1)
	}

	configCasbinEnforcer()
	configRedisClient()
	configInitCacheStore()
}

