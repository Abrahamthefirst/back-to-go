package entities

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
type AccesTokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
