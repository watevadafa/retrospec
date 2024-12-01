package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // "-" means this won't be included in JSON
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Plan         string    `json:"plan"` // "free" or "pro"
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewUser(email, first_name, last_name, plan string) *User {
	now := time.Now()
	return &User{
		ID:        uuid.New().String(),
		Email:     email,
		FirstName: first_name,
		LastName:  last_name,
		Plan:      plan,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
