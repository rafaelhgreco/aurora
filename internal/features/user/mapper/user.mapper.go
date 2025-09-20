package mapper

import (
	"aurora.com/aurora-backend/internal/features/user/domain"
	"aurora.com/aurora-backend/internal/features/user/dto"
)

func FromCreateRequestToUserEntity(req *dto.CreateUserRequest) (*domain.User, error) {
	userType, err := mapStringToUserType(req.Type)
	if err != nil {
		return nil, err
	}

	adminData, collaboratorData := buildRoleSpecificData(userType, req)
	return &domain.User{
		Name:             req.Name,
		Email:            req.Email,
		Password:         req.Password,
		Type:             userType,
		AdminData:        adminData,
		CollaboratorData: collaboratorData,
	}, nil
}

func FromUserEntityToUserResponse(entity *domain.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		CreatedAt: entity.CreatedAt,
	}
}

func FromUserEntityToSpecificResponse(entity *domain.User) interface{} {
	base := dto.UserResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Type:      entity.Type.String(),
		CreatedAt: entity.CreatedAt,
	}

	switch entity.Type {
	case domain.ADMIN:
		return &dto.AdminUserResponse{
			UserResponse: base,
			Permissions:      entity.AdminData.Permissions,
		}
	case domain.COLLABORATOR:
		return &dto.CollaboratorUserResponse{
			UserResponse: base,
			TeamID:           entity.CollaboratorData.TeamID,
		}
	}
	return &base
}