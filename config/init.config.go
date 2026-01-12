package config

func InitConfig() {
	configLoadEnv()
	configJwtKey()
	configDatabaseConnection()
	configCasbinEnforcer()
	configRedisClient()
	configCSRF()
	configInitCacheStore()
}
