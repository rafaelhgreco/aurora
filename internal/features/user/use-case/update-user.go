package usecase

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

type UpdateUserUseCase struct {
	userRepo domain.UserRepository
}

func NewUpdateUserUseCase(repo domain.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepo: repo}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, id string, user *domain.User) (*domain.User, error) {

	existingUser, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	existingUser.Name = user.Name
	existingUser.Email = user.Email

	return uc.userRepo.Update(ctx, existingUser)
}
