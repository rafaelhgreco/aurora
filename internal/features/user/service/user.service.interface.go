package service

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/dto"
	securityDTO "aurora.com/aurora-backend/internal/features/user/security/dto"
)

type IUserService interface {
    // --- Métodos CRUD ---
    Create(ctx context.Context, req *dto.CreateUserRequest) (interface{}, error) 
    
    // Sugestão: Mudar para interface{} para suportar respostas polimórficas.
    FindByID(ctx context.Context, id string) (interface{}, error)
    
    // Sugestão: Mudar para interface{} para suportar respostas polimórficas.
    Update(ctx context.Context, id string, req *dto.UpdateUserRequest) (interface{}, error)
    
    Delete(ctx context.Context, id string) error

    // --- Métodos de Segurança ---
    Login(ctx context.Context, idToken string) (*securityDTO.LoginResponse, error)
    ChangePassword(ctx context.Context, userID string, req *securityDTO.ChangePasswordRequest) error
}