package authserivce

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Abrahamthefirst/back-to-go/internal/dtos"
	"github.com/Abrahamthefirst/back-to-go/internal/entities"
	"github.com/Abrahamthefirst/back-to-go/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) SignUp(dto dtos.SignupRequestDto) (*entities.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	user := entities.User{
		Email:    dto.Email,
		Password: string(hashedPassword),
		Username: dto.Username,
	}
	newUser, err := s.repo.Create(user)

	if err != nil {
		return nil, fmt.Errorf("%w: user with email %s already exists", entities.ErrConflict, dto.Email)
	}

	return newUser, nil
}

type LoginServiceOkResponse struct {
	User        *entities.User `json:"user"`
	AccessToken string         `json:"access_token"`
}

func (a *AuthService) Login(ctx context.Context, dto dtos.LoginRequestDto) (*LoginServiceOkResponse, error) {
	user, err := a.repo.FindByEmail(ctx, dto.Email)
	if err != nil {
		if errors.Is(err, entities.RecordNotFound) {
			return nil, entities.ErrInvalidCredentials

		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return nil, entities.ErrInvalidCredentials
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "my-auth-server",
		"id":  user.ID,
	})

	token, err := claims.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return nil, entities.ErrInternal
	}

	return &LoginServiceOkResponse{User: user, AccessToken: token}, nil

}
