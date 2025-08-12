package usecase

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

type CreateUserUseCase struct {
	userRepo domain.UserRepository
}

func NewCreateUserUseCase(repo domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: repo}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, user *domain.User) (*domain.User, error) {
	return uc.userRepo.Save(ctx, user)
}