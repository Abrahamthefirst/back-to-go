package dtos

import "github.com/Abrahamthefirst/back-to-go/internal/entities"

type LoginResponseBody struct {
	AccessToken string        `json:"access_token"`
	User        entities.User `json:"user"`
}
type SignupResponseBody struct {
	User entities.User `json:"user"`
}
type LoginResponse struct {
	Message    string            `json:"message"`
	Data       LoginResponseBody `json:"data"`
	StatusCode uint              `json:"statusCode"`
}

type ErrorResponse struct {
	Message    string `json:"message"`
	Error      any    `json:"error"`
	StatusCode uint   `json:"statusCode"`
}

type SignupResponse struct {
	Message    string             `json:"message"`
	Data       SignupResponseBody `json:"data"`
	StatusCode uint               `json:"statusCode"`
}
type LoginRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=4,lte=20"`
}

type SignupRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=4,lte=20"`
	Username string `json:"username" validate:"required,gte=3,lte=20"`
}
