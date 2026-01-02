package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL         string
	PORT                 string
	ACCESS_TOKEN_SECRET  string
	REFRESH_TOKEN_SECRET string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No env file found, using system env")
	}

	return &Config{
		DATABASE_URL:         GetEnv("DATABASE_URL", "postgresql://postgres:abraham@localhost:5432/back-to-go"),
		PORT:                 GetEnv("PORT", ":4000"),
		ACCESS_TOKEN_SECRET:  GetEnv("ACCESS_TOKEN_SECRET", ""),
		REFRESH_TOKEN_SECRET: GetEnv("REFRESH_TOKEN_SECRET", ""),
	}
}

func GetEnv(key string, fallback string) string {
	val, exist := os.LookupEnv(key)

	if !exist {
		log.Printf("%v is not set", key)
		return fallback
	}
	return val
}
