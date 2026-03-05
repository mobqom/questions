package config

import "os"

type AppConfig struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
	Port   string
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func Init() *AppConfig {
	return &AppConfig{
		DbHost: getEnv("DB_HOST", "localhost"),
		DbPort: getEnv("DB_PORT", "5432"),
		DbUser: getEnv("DB_USER", "user"),
		DbPass: getEnv("DB_PASS", "password"),
		DbName: getEnv("DB_NAME", "questions"),
		Port:   getEnv("PORT", "8081"),
	}
}
