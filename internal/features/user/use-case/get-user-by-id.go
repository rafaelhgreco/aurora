package usecase

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

type GetUserByIDUseCase struct {
	userRepo domain.UserRepository
}

func NewGetUserByIDUseCase(repo domain.UserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{userRepo: repo}
}

func (uc *GetUserByIDUseCase) Execute(ctx context.Context, id int) (*domain.User, error) {
	return uc.userRepo.FindByID(ctx, id)
}
