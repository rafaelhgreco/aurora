package mapper

import (
	"fmt"

	"aurora.com/aurora-backend/internal/features/user/domain"
	"aurora.com/aurora-backend/internal/features/user/dto"
)

func mapStringToUserType(typeStr string) (domain.UserType, error) {
	switch typeStr {
	case "ADMIN":
		return domain.ADMIN, nil
	case "COLLABORATOR":
		return domain.COLLABORATOR, nil
	case "FINANCIAL":
		return domain.FINANCIAL, nil
	case "COMMON":
		return domain.COMMON, nil
	default:
		return 0, fmt.Errorf("invalid user type: %s", typeStr)
	}
}

func buildRoleSpecificData(userType domain.UserType, req *dto.CreateUserRequest) (*domain.AdminProfile, *domain.CollaboratorProfile, *domain.FinancialProfile) {
	switch userType {
	case domain.ADMIN:
		return &domain.AdminProfile{Permissions: req.Permissions}, nil, nil
	case domain.COLLABORATOR:
		return nil, &domain.CollaboratorProfile{TeamID: req.TeamID}, nil
	case domain.FINANCIAL:
		financialProfile := mapFinancialProfileFromDTO(req)
		return nil, nil, financialProfile
	default:
		return nil, nil, nil
	}
}