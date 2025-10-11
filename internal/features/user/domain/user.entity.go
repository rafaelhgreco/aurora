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
	return [...]string{"COMMON", "COLLABORATOR", "ADMIN", ""}[ut]
}

type AdminProfile struct {
	Permissions []string
}

type CollaboratorProfile struct {
	TeamID string
	Projects []string
}

type AuthClient interface {
	VerifyIDToken(ctx context.Context, idToken string) (string, error)
	GenerateAccessToken(ctx context.Context, userID string) (string, error)
	GenerateRefreshToken(ctx context.Context, userID string) (string, error)
	UpdateUser(ctx context.Context, uid string, params interface{}) (interface{}, error)
	CreateUser(ctx context.Context, user *User) (string, error) 
}

type ChangePasswordInput struct {
	UserID string
	NewPassword string
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

type ErrUserNotFound struct {
	ID string
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("user with id %s not found", e.ID)
}