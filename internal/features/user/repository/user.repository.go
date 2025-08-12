package repository

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

// UserRepository agora inclui métodos para atualizar e deletar.
type UserRepository interface {
	Save(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByID(ctx context.Context, id int) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error // <--- NOVO
	Delete(ctx context.Context, id int) error          // <--- NOVO
}