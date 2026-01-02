package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	handler "github.com/Abrahamthefirst/back-to-go/internal/api"
	"github.com/Abrahamthefirst/back-to-go/internal/config"
	"github.com/Abrahamthefirst/back-to-go/internal/repository"
	authserivce "github.com/Abrahamthefirst/back-to-go/internal/service"
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

	a.router.GET("/health", func(c *gin.Context) {
		slog.WarnContext(c.Request.Context(), "request validation failed",
			"path", c.Request.URL.Path,
		)
		c.JSON(http.StatusOK, gin.H{
			"message": "Health-Check",
		})
	})

	timeoutMiddleware := func(timeout time.Duration) gin.HandlerFunc {
		return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
			defer cancel()

			c.Request = c.Request.WithContext(ctx)
			c.Next()
		}
	}

	authMiddleware := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			header := c.GetHeader("Authorization")
			a.logger.Log(c.Request.Context(), slog.LevelInfo, header)

			token := strings.Split(header, "Bearer ")
			fmt.Println(token[1])

			c.Next()
		}
	}

	a.router.Use(timeoutMiddleware(time.Second * 2))
	a.router.Use(authMiddleware())

	UserRepository := repository.NewUserRepository(a.db)
	authserivce := authserivce.NewAuthService(UserRepository)
	authController := handler.NewAuthController(authserivce)

	handler.RegisterAuthRoutes(a.router, authController)

	a.router.Run(":3000")

}
