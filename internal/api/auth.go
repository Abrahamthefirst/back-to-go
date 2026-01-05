package handler

import (
	"log/slog"
	"net/http"

	"github.com/Abrahamthefirst/back-to-go/internal/dtos"
	authserivce "github.com/Abrahamthefirst/back-to-go/internal/service/auth-service"

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
		webutil.HandleError(c, err)
		return
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

func RegisterAuthRoutes(r *gin.Engine, authController *AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Signup)
		auth.POST("/login", authController.Login)
		auth.POST("/usera", authController.Login)
	}
}
