package service

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/dto"
	securityDTO "aurora.com/aurora-backend/internal/features/user/security/dto"
)

type IUserService interface {
    Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) 
    
    FindByID(ctx context.Context, id string) (*dto.UserResponse, error)
    
    Update(ctx context.Context, id string, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
    
    Delete(ctx context.Context, id string) error
    

    Login(ctx context.Context, idToken string) (*securityDTO.LoginResponse, error)
    ChangePassword(ctx context.Context, userID string, req *securityDTO.ChangePasswordRequest) error
}