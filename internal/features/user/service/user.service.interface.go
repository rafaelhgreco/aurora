package service

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/dto"
)

type IUserService interface {
	Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error)
	FindByID(ctx context.Context, id string) (*dto.UserResponse, error)
	Update(ctx context.Context, id string, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(ctx context.Context, id string) error
}