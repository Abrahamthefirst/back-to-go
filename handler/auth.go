package handler

import (
	"log/slog"
	"net/http"

	authserivce "github.com/Abrahamthefirst/back-to-go/auth-serivce"
	"github.com/Abrahamthefirst/back-to-go/webutil"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	logger      *slog.Logger
	authService *authserivce.AuthService
}

func NewAuthController(authSerivce *authserivce.AuthService) *AuthController {
	return &AuthController{
		logger:      slog.Default().With("component", "auth"),
		authService: authSerivce,
	}
}

func (a *AuthController) Login(c *gin.Context) {

	var requestDto LoginRequestDto

	if err := webutil.ValidateRequest(c, &requestDto); err != nil {
		return
	}

	response := LoginResponse{
		Message: "Login Successfull",
		Body:    LoginResponseBody{AccessToken: token},
	}

	c.JSON(http.StatusOK, response)
}

func (a *AuthController) Signup(c *gin.Context) {

}
func (a *AuthController) ChangePassword(c *gin.Context) {}

func RegisterAuthRoutes(r *gin.Engine, authController *AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Signup)
	}
}
