package application

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Abrahamthefirst/back-to-go/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type application struct {
	router *gin.Engine
}

func NewApp() *application {
	router := gin.Default()

	return &application{
		router: router,
		
	}

}

func (a *application) Bootstrap() {
	err := godotenv.Load()
	if err != nil {
		panic("")
	}

	handlerOptions := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}

	handler := slog.NewJSONHandler(os.Stdout, handlerOptions)

	logger := slog.New(handler)

	slog.SetDefault(logger)
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

	authController := handler.NewAuthController()
	
	handler.RegisterAuthRoutes(a.router, authController)



	a.router.Run(":3000")
}
