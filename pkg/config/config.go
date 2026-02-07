package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		Port:       getEnv("PORT", "8000"),
		DBHost:     getEnv("db_host", "localhost"),
		DBUser:     getEnv("db_user", "postgres"),
		DBPassword: getEnv("db_pass", ""),
		DBName:     getEnv("db_name", "businessense"),
		JWTSecret:  getEnv("token_password", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
