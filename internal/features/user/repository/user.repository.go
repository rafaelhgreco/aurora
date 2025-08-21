package repository

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}