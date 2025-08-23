package domain

import (
	"context"
	"fmt"
	"time"
)

type UserType int

const (
	COMMON UserType = iota
	COLLABORATOR
	ADMIN
)

func (ut UserType) String() string {
	return [...]string{"COMMON", "COLLABORATOR", "ADMIN"}[ut]
}

type AdminProfile struct {
	Permissions []string
}

type CollaboratorProfile struct {
	TeamID string
	Projects []string
}


type User struct {
	ID        string 	`firestore:"-"`
	Name      string    `firestore:"name"`
	Email     string    `firestore:"email"`
	Password  string    `firestore:"password"`
	CreatedAt time.Time `firestore:"createdAt"`

	Type      UserType  `firestore:"type"`
	AdminData *AdminProfile `firestore:"adminData,omitempty"`
	CollaboratorData *CollaboratorProfile `firestore:"collaboratorData,omitempty"`
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