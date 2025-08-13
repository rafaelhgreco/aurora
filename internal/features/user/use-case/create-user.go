package usecase

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

type CreateUserUseCase struct {
	userRepo domain.UserRepository
	hasher domain.PasswordHasher
}

func NewCreateUserUseCase(repo domain.UserRepository, hasher domain.PasswordHasher) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: repo, hasher: hasher}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashed, err := uc.hasher.Hash(ctx, user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashed
	return uc.userRepo.Save(ctx, user)
}