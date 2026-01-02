package main

import (
	"github.com/Abrahamthefirst/back-to-go/internal/config"
	"github.com/Abrahamthefirst/back-to-go/internal/database"
	"github.com/Abrahamthefirst/back-to-go/pkg/logger"
)

func main() {
	cfg := config.Load()
	db := database.NewPgDB(cfg.DATABASE_URL)

	app := NewApp(db, cfg, logger.New(false))

	app.Bootstrap()

}
