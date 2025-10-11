package mapper

import (
	"aurora.com/aurora-backend/internal/features/user/domain"
	domainDTO "aurora.com/aurora-backend/internal/features/user/dto"
	securityDTO "aurora.com/aurora-backend/internal/features/user/security/dto"
)

func FromCreateRequestToUserEntity(req *domainDTO.CreateUserRequest) (*domain.User, error) {
	userType, err := mapStringToUserType(req.Type)
	if err != nil {
		return nil, err
	}
	if req.Type == "" {
		userType, err = mapStringToUserType(req.Type)
		if err != nil {
			return nil, err
		}
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

func FromUserEntityToUserResponse(entity *domain.User) *domainDTO.UserResponse {
	return &domainDTO.UserResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Type:      entity.Type.String(),
		CreatedAt: entity.CreatedAt,
	}
}

func FromUserEntityToSpecificResponse(entity *domain.User) interface{} {
	base := domainDTO.UserResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Type:      entity.Type.String(),
		CreatedAt: entity.CreatedAt,
	}

	switch entity.Type {
	case domain.ADMIN:
		return &domainDTO.AdminUserResponse{
			UserResponse: base,
			Permissions:      entity.AdminData.Permissions,
		}
	case domain.COLLABORATOR:
		return &domainDTO.CollaboratorUserResponse{
			UserResponse: base,
			TeamID:           entity.CollaboratorData.TeamID,
		}
	}
	return &base
}

func FromUpdateRequestToUserEntity(req *domainDTO.UpdateUserRequest) (*domain.User, error) {
    return &domain.User{
        Name:  derefString(req.Name),
        Email: derefString(req.Email),
    }, nil
}

func derefString(s *string) string {
    if s == nil {
        return ""
    }
    return *s
}

func FromChangePasswordRequestToDomain(userID string, req *securityDTO.ChangePasswordRequest) *domain.ChangePasswordInput {
    return &domain.ChangePasswordInput{
        UserID:      userID,
        NewPassword: req.NewPassword,
    }
}