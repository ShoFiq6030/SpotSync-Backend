package config

import (
	
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	Dsn       string
	JwtSecret string
}

func LoadEnv() *Config {
	// Ignore the error if .env doesn't exist
	_ = godotenv.Load()

	return &Config{
		Port:      getEnv("PORT", "8080"),
		Dsn:       getEnv("DSN", ""),
		JwtSecret: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}