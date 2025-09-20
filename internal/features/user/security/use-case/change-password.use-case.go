package usecase

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
	"firebase.google.com/go/v4/auth"
)

type ChangePasswordUseCase struct {
	authClient domain.AuthClient
}

func NewChangePasswordUseCase(authClient domain.AuthClient) *ChangePasswordUseCase {
	return &ChangePasswordUseCase{authClient: authClient}
}

func (uc *ChangePasswordUseCase) Execute(ctx context.Context, uid string, newPassword string) error {
	params := (&auth.UserToUpdate{}).Password(newPassword)
	_, err := uc.authClient.UpdateUser(ctx, uid, params)
	return err
}