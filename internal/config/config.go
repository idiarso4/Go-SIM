package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	ServerPort     int
	Environment    string
	JWTSecret      string
	JWTExpiration  time.Duration
	UploadDir      string
	ExportDir      string
}

func LoadConfig() *Config {
	port, _ := strconv.Atoi(getEnvOrDefault("PORT", "8080"))
	jwtExpStr := getEnvOrDefault("JWT_EXPIRATION", "24h")
	jwtExp, _ := time.ParseDuration(jwtExpStr)

	return &Config{
		ServerPort:     port,
		Environment:    getEnvOrDefault("ENV", "development"),
		JWTSecret:      getEnvOrDefault("JWT_SECRET", "your_secret_key_here"),
		JWTExpiration:  jwtExp,
		UploadDir:      getEnvOrDefault("UPLOAD_DIR", "./uploads"),
		ExportDir:      getEnvOrDefault("EXPORT_DIR", "./exports"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}