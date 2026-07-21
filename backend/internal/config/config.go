package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config stores all application configuration.
type Config struct {
	// Application
	AppName string
	AppEnv  string
	AppPort string

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// JWT
	JWTSecret               string
	JWTAccessTokenDuration  time.Duration
	JWTRefreshTokenDuration time.Duration
}

// Load loads configuration from environment variables.
func Load() *Config {
	// Ignore error if .env does not exist.
	_ = godotenv.Load()

	cfg := &Config{
		AppName: getEnv("APP_NAME", "jck-connect"),
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "jck_connect"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		JWTSecret: getEnv(
			"JWT_SECRET",
			"change-this-secret-before-production",
		),

		JWTAccessTokenDuration: parseDuration(
			getEnv("JWT_ACCESS_TOKEN_DURATION", "15m"),
			15*time.Minute,
		),

		JWTRefreshTokenDuration: parseDuration(
			getEnv("JWT_REFRESH_TOKEN_DURATION", "168h"),
			168*time.Hour,
		),
	}

	log.Printf(
		"Configuration loaded (env=%s, port=%s)",
		cfg.AppEnv,
		cfg.AppPort,
	)

	return cfg
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func parseDuration(value string, fallback time.Duration) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return duration
}
