package authserivce

import (
	"os"

	"github.com/Abrahamthefirst/back-to-go/repository"
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

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginServiceOkResponse struct {
	AccessToken string `json:"access_token"`
	user        *repository.User
}

func (a *AuthService) Login(dto LoginRequestDto) (*LoginServiceOkResponse, error) {
	user, err := a.repo.FindByEmail(dto.Email)

	if err != nil {

	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return nil, err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "my-auth-server",
		"id":  user.ID,
	})

	token, err := claims.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return
	}

	return LoginServiceOkResponse{user: user, AccessToken: token}

}

func (a *AuthService) SignUp(dto LoginRequestDto) {
	user, err := a.repo.Create(dto)

	if err != nil {
		return
	}
}
