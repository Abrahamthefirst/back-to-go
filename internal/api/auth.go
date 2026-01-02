package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Abrahamthefirst/back-to-go/internal/dtos"
	"github.com/Abrahamthefirst/back-to-go/internal/entities"
	authserivce "github.com/Abrahamthefirst/back-to-go/internal/service"
	"github.com/Abrahamthefirst/back-to-go/internal/webutil"
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

func (a *AuthController) Signup(c *gin.Context) {

	var requestDto dtos.SignupRequestDto

	if err := webutil.ValidateRequest(c, &requestDto); err != nil {
		return
	}

	user, err := a.authService.SignUp(requestDto)
	if err != nil {
		if errors.Is(err, entities.ErrConflict) {
			c.JSON(http.StatusConflict, dtos.ErrorResponse{Message: err.Error(), Error: err, StatusCode: http.StatusConflict})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})

	}

	response := dtos.SignupResponse{
		Message:    "Registration Successful",
		Data:       dtos.SignupResponseBody{User: *user},
		StatusCode: 201,
	}

	c.JSON(http.StatusOK, response)
}

func (a *AuthController) Login(c *gin.Context) {

	var requestDto dtos.LoginRequestDto

	if err := webutil.ValidateRequest(c, &requestDto); err != nil {
		return
	}

	result, err := a.authService.Login(c.Request.Context(), requestDto)
	if err != nil {
		webutil.HandleError(c, err)
		return
	}

	response := dtos.LoginResponse{
		Message:    "Login Successfull",
		Data:       dtos.LoginResponseBody{User: *result.User, AccessToken: result.AccessToken},
		StatusCode: 200,
	}

	c.JSON(http.StatusOK, response)

}

func (a *AuthController) ChangePassword(c *gin.Context) {}

func RegisterAuthRoutes(r *gin.Engine, authController *AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Signup)
		auth.POST("/login", authController.Login)
	}
}
