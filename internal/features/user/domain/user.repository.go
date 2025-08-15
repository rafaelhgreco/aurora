package domain

import "context"

type UserRepository interface {
	Save(ctx context.Context, user *User) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
}