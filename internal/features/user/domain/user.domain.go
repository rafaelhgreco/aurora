package domain

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	ID        string 	`firestore:"-"`
	Name      string    `firestore:"name"`
	Email     string    `firestore:"email"`
	Password  string    `firestore:"password"`
	CreatedAt time.Time `firestore:"createdAt"`
}

type PasswordHasher interface {
    Hash(ctx context.Context, plain string) (string, error)
}

type ErrUserNotFound struct {
	ID string
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("user with id %s not found", e.ID)
}