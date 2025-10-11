package usecase

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

type CreateUserUseCase struct {
	userRepo   domain.UserRepository
	authClient domain.AuthClient
}

func NewCreateUserUseCase(repo domain.UserRepository, authClient domain.AuthClient) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: repo, authClient: authClient}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, user *domain.User) (*domain.User, error) {
	uid, err := uc.authClient.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = uid
	user.Password = ""

	savedUser, err := uc.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}
