package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv          string
	AppPort         string
	DatabaseURL     string
	DefaultLanguage string
	AutoMigrate     bool
	SeedDB          bool
	DBDebug         bool
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		AppEnv:          getEnv("APP_ENV", "development"),
		AppPort:         getEnv("APP_PORT", "3011"),
		DatabaseURL:     getEnv("DATABASE_URL", ""),
		DefaultLanguage: getEnv("DEFAULT_LANGUAGE", "en"),
		AutoMigrate:     getEnv("AUTO_MIGRATE", "true") == "true",
		SeedDB:          getEnv("SEED_DB", "false") == "true",
		DBDebug:         getEnv("DB_DEBUG", "false") == "true",
	}
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}
