package config

func InitConfig() {
	configLoadEnv()
	configJwtKey()
	configDatabaseConnection()
	configCasbinEnforcer()
}
