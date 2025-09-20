package mapper

import (
	"aurora.com/aurora-backend/internal/features/user/domain"
	"aurora.com/aurora-backend/internal/features/user/dto"
)

func mapStringToUserType(typeStr string) (domain.UserType, error) {
	switch typeStr {
	case "ADMIN":
		return domain.ADMIN, nil
	case "COLLABORATOR":
		return domain.COLLABORATOR, nil
	default:
		return domain.COMMON, nil
	}
}

func buildRoleSpecificData(userType domain.UserType, req *dto.CreateUserRequest) (*domain.AdminProfile, *domain.CollaboratorProfile) {
	switch userType {
	case domain.ADMIN:
		return &domain.AdminProfile{Permissions: req.Permissions}, nil
	case domain.COLLABORATOR:
		return nil, &domain.CollaboratorProfile{TeamID: req.TeamID}
	default:
		return nil, nil
	}
}