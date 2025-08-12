package service

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/dto"
	"aurora.com/aurora-backend/internal/features/user/mapper"
	usecase "aurora.com/aurora-backend/internal/features/user/use-case"
)

type userService struct {
	createUserUseCase *usecase.CreateUserUseCase
	getUserUseCase    *usecase.GetUserByIDUseCase
}

func NewUserService(
	createUserUC *usecase.CreateUserUseCase,
	getUserUC *usecase.GetUserByIDUseCase,
) IUserService {
	return &userService{
		createUserUseCase: createUserUC,
		getUserUseCase:    getUserUC,
	}
}
func (s *userService) Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	userEntity := mapper.FromCreateRequestToUserEntity(req)

	createdUser, err := s.createUserUseCase.Execute(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	response := mapper.FromUserEntityToUserResponse(createdUser)
	return response, nil
}

func (s *userService) FindByID(ctx context.Context, id int) (*dto.UserResponse, error) {
	user, err := s.getUserUseCase.Execute(ctx, id)
	if err != nil {
		return nil, err
	}

	response := mapper.FromUserEntityToUserResponse(user)
	return response, nil
}