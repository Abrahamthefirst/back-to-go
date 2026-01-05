package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	handler "github.com/Abrahamthefirst/back-to-go/internal/api"
	"github.com/Abrahamthefirst/back-to-go/internal/config"
	"github.com/Abrahamthefirst/back-to-go/internal/middleware"
	"github.com/Abrahamthefirst/back-to-go/internal/repository"
	authserivce "github.com/Abrahamthefirst/back-to-go/internal/service/auth-service"
	emailservice "github.com/Abrahamthefirst/back-to-go/internal/service/email-service"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type application struct {
	router *gin.Engine
	db     *gorm.DB
	cfg    *config.Config
	logger *slog.Logger
}

func NewApp(db *gorm.DB, cfg *config.Config, logger *slog.Logger) *application {
	router := gin.Default()

	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	}
	return &application{
		router: router,
		db:     db,
		cfg:    cfg,
		logger: logger,
	}

}

func (a *application) Bootstrap() {

	slog.SetDefault(a.logger)

	a.startServer()

}

func (a *application) startServer() {

	a.router.Use(middleware.AuthMiddleware())

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))

	if err != nil {
		slog.Warn("Could not convert smtp port to integer")
	}

	email, err := emailservice.New(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_USER"))

	if err != nil {
		slog.Warn(fmt.Sprintf("Failed to initialize email service: %v", err))
	}

	userRepository := repository.NewUserRepository(a.db)

	authserivce := authserivce.NewAuthService(userRepository, email)

	authController := handler.NewAuthController(authserivce)

	handler.RegisterAuthRoutes(a.router, authController)

	a.router.Run(a.cfg.PORT)

}
