package repository

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByID(ctx context.Context, id int) (*domain.User, error)
}