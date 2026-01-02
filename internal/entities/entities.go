package entities

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
