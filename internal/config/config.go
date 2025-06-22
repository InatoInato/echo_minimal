package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBHost  string
	DBUser  string
	DBPass  string
	DBName  string
	DBPort  string
}

var Cfg *Config

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Cannot to lead .env file")
	}

	Cfg = &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBUser:  getEnv("DB_USER", "postgres"),
		DBPass:  getEnv("DB_PASSWORD", "postgres"),
		DBName:  getEnv("DB_NAME", "echo"),
		DBPort:  getEnv("DB_PORT", "5432"),
	}
}
