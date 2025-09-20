package usecase

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

type DeleteUserUseCase struct {
	userRepo domain.UserRepository
}

func NewDeleteUserUseCase(repo domain.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{userRepo: repo}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, id string) error {
	return uc.userRepo.Delete(ctx, id)
}
