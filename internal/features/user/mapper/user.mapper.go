package mapper

import (
	"aurora.com/aurora-backend/internal/features/user/domain"
	"aurora.com/aurora-backend/internal/features/user/dto"
)

func FromCreateRequestToUserEntity(req *dto.CreateUserRequest) *domain.User {
	return &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, //o hash seria feito no use-case
	}
}

func FromUserEntityToUserResponse(entity *domain.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		CreatedAt: entity.CreatedAt,
	}
}