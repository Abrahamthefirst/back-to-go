package database

import (
	"time"

	"github.com/Abrahamthefirst/back-to-go/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
}

func NewPgDB(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = db.AutoMigrate(
		&repository.UserModel{},
	)
	if err != nil {
		panic("Database connection failed")
	}

	return db

}
