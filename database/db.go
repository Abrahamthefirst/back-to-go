package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
}

func NewPgDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection failed")
	}

	return db

}
