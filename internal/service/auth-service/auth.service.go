package authserivce

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Abrahamthefirst/back-to-go/internal/dtos"
	"github.com/Abrahamthefirst/back-to-go/internal/entities"
	"github.com/Abrahamthefirst/back-to-go/internal/repository"
	emailservice "github.com/Abrahamthefirst/back-to-go/internal/service/email-service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wneessen/go-mail"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo         *repository.UserRepository
	emailService *emailservice.Mailer
}

func NewAuthService(repo *repository.UserRepository, email *emailservice.Mailer) *AuthService {
	return &AuthService{
		repo:         repo,
		emailService: email,
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

	}
	message := "Welcome to vayla"

	go func() {
		s.emailService.SendEmail(&emailservice.MailConfig{To: "abrahamodusegun36@gmail.com", Format: mail.TypeTextPlain, Message: &message})
	}()

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

	claims := entities.AccesTokenClaims{
		UserID: strconv.Itoa(int(user.ID)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my-auth-server",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))

	if err != nil {
		return nil, entities.ErrInternal
	}

	return &LoginServiceOkResponse{User: user, AccessToken: tokenString}, nil

}

func (a *AuthService) GetUsers(ctx context.Context) (*[]entities.User, error) {

	users, err := a.repo.GetAllUsers(ctx)

	if err != nil {
		return nil, err
	}
	return users, nil

}
func (a *AuthService) ChangePassword(ctx context.Context, dto dtos.ChanePasswordRequestDto) {

}
