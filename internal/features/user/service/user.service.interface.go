package service

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/dto"
)

type IUserService interface {
	Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error)
	FindByID(ctx context.Context, id string) (*dto.UserResponse, error)
}