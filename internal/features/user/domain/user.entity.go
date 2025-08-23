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
	FINANCIAL
)

func (ut UserType) String() string {
	return [...]string{"COMMON", "COLLABORATOR", "ADMIN", "FINANCIAL"}[ut]
}

type AdminProfile struct {
	Permissions []string
}

type CollaboratorProfile struct {
	TeamID string
	Projects []string
}

type BillingAddress struct {
	Street  string
	City   string
	State   string
	ZipCode string
	Country string
}

type Subscription struct {
	PlanID      string
	StartDate time.Time
	EndDate   time.Time
	Active    bool
}

type FinancialProfile struct {
	BillingAddress BillingAddress
	Subscription   Subscription
	PaymentMethods string
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
	FinancialData *FinancialProfile `firestore:"financialData,omitempty"`
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