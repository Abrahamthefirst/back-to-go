package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL         string
	PORT                 string
	ACCESS_TOKEN_SECRET  string
	REFRESH_TOKEN_SECRET string
	SMTP_USERNAME        string
	SMTP_PASSWORD        string
}

func Load() *Config {
	if err := godotenv.Load("../../.env"); err != nil {
		slog.Warn("No env file found, using system env")
	}

	return &Config{
		DATABASE_URL:         GetEnv("DATABASE_URL", "postgresql://postgres:abraham@localhost:5432/back-to-go"),
		PORT:                 GetEnv("PORT", ":4000"),
		ACCESS_TOKEN_SECRET:  GetEnv("ACCESS_TOKEN_SECRET", ""),
		REFRESH_TOKEN_SECRET: GetEnv("REFRESH_TOKEN_SECRET", ""),
		SMTP_USERNAME:        GetEnv("SMTP_USERNAME", ""),
		SMTP_PASSWORD:        GetEnv("SMTP_PASSWORD", ""),
	}
}

func GetEnv(key string, fallback string) string {
	val, exist := os.LookupEnv(key)

	if !exist {
		slog.Warn(fmt.Sprintf("%v is not set", key))
		return fallback
	}
	return val
}
