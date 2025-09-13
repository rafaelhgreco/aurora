package usecase

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

type LoginUserUseCase struct {
	userRepo domain.UserRepository
	authClient domain.AuthClient
}

func NewLoginUserUseCase(userRepo domain.UserRepository, authClient domain.AuthClient) *LoginUserUseCase {
	return &LoginUserUseCase{
		userRepo: userRepo,
		authClient: authClient,
	}
}

func (uc *LoginUserUseCase) Execute(ctx context.Context, idToken string) (string, string, *domain.User, error) {
	// Verify the ID token with the external auth provider
	userID, err := uc.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", "", nil, err
	}

	// Retrieve user from the database
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return "", "", nil, err
	}

	// Generate access and refresh tokens
	accessToken, err := uc.authClient.GenerateAccessToken(ctx, user.ID)
	if err != nil {
		return "", "", nil, err
	}

	refreshToken, err := uc.authClient.GenerateRefreshToken(ctx, user.ID)
	if err != nil {
		return "", "", nil, err
	}

	return accessToken, refreshToken, user, nil
}