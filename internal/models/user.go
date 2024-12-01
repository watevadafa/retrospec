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

func NewUser(email, first_name string, last_name *string, plan ...string) *User {
	now := time.Now()

	userPlan := "free" // Default plan
	if len(plan) > 0 && plan[0] != "" {
		userPlan = plan[0]
	}

	return &User{
		ID:        uuid.New().String(),
		Email:     email,
		FirstName: first_name,
		LastName:  *last_name,
		Plan:      userPlan,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
