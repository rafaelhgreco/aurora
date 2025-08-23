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
	updateUserUseCase *usecase.UpdateUserUseCase
	deleteUserUseCase *usecase.DeleteUserUseCase
}

func NewUserService(
	createUserUC *usecase.CreateUserUseCase,
	getUserUC *usecase.GetUserByIDUseCase,
	updateUserUC *usecase.UpdateUserUseCase,
	deleteUserUC *usecase.DeleteUserUseCase,
) IUserService {
	return &userService{
		createUserUseCase: createUserUC,
		getUserUseCase:    getUserUC,
		updateUserUseCase: updateUserUC,
		deleteUserUseCase: deleteUserUC,
	}
}
func (s *userService) Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	userEntity, err := mapper.FromCreateRequestToUserEntity(req)
	if err != nil {
		return nil, err
	}
	
	createdUser, err := s.createUserUseCase.Execute(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	response := mapper.FromUserEntityToUserResponse(createdUser)
	return response, nil
}

func (s *userService) FindByID(ctx context.Context, id string) (*dto.UserResponse, error) {
	user, err := s.getUserUseCase.Execute(ctx, id)
	if err != nil {
		return nil, err
	}

	response := mapper.FromUserEntityToUserResponse(user)
	return response, nil
}

func (s *userService) Update(ctx context.Context, id string, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	updatedUser, err := s.updateUserUseCase.Execute(ctx, id, req.Name, req.Email)
	if err != nil {
		return nil, err
	}

	response := mapper.FromUserEntityToUserResponse(updatedUser)
	return response, nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	return s.deleteUserUseCase.Execute(ctx, id)
}