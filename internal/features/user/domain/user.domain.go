package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type PasswordHasher interface {
    Hash(ctx context.Context, plain string) (string, error)
}