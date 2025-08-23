package service

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/dto"
	"aurora.com/aurora-backend/internal/features/user/mapper"
	securityDTO "aurora.com/aurora-backend/internal/features/user/security/dto"
	usecaseSecurity "aurora.com/aurora-backend/internal/features/user/security/use-case"
	usecase "aurora.com/aurora-backend/internal/features/user/use-case"
)

type userService struct {
	createUserUseCase *usecase.CreateUserUseCase
	getUserUseCase    *usecase.GetUserByIDUseCase
	updateUserUseCase *usecase.UpdateUserUseCase
	deleteUserUseCase *usecase.DeleteUserUseCase
	loginUseCase *usecaseSecurity.LoginUserUseCase
	changePasswordUC *usecaseSecurity.ChangePasswordUseCase
}

func NewUserService(
	createUserUC *usecase.CreateUserUseCase,
	getUserUC *usecase.GetUserByIDUseCase,
	updateUserUC *usecase.UpdateUserUseCase,
	deleteUserUC *usecase.DeleteUserUseCase,
	loginUC *usecaseSecurity.LoginUserUseCase,
	changeUC *usecaseSecurity.ChangePasswordUseCase,
) IUserService {
	return &userService{
		createUserUseCase: createUserUC,
		getUserUseCase:    getUserUC,
		updateUserUseCase: updateUserUC,
		deleteUserUseCase: deleteUserUC,
		loginUseCase: loginUC,
		changePasswordUC: changeUC,
	}
}
func (s *userService) Create(ctx context.Context, req *dto.CreateUserRequest) (interface{}, error) {
	// 1. Usa o Mapper para traduzir o DTO em uma Entidade de Domínio.
	userEntity, err := mapper.FromCreateRequestToUserEntity(req)
	if err != nil {
		// Retorna um erro se o mapeamento falhar (ex: tipo de usuário inválido)
		return nil, err
	}

	// 2. AGORA sim, passa a Entidade ('userEntity') para o UseCase.
	createdUser, err := s.createUserUseCase.Execute(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	// 3. Usa o outro lado do Mapper para traduzir a Entidade de volta para um DTO de resposta.
	response := mapper.FromUserEntityToSpecificResponse(createdUser)
	return response, nil
}

func (s *userService) FindByID(ctx context.Context, id string) (interface{}, error) {
	user, err := s.getUserUseCase.Execute(ctx, id)
	if err != nil {
		return nil, err
	}

	response := mapper.FromUserEntityToUserResponse(user)
	return response, nil
}

func (s *userService) Update(ctx context.Context, id string, req *dto.UpdateUserRequest) (interface{}, error) {
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

func (s *userService) Login(ctx context.Context, idToken string) (*securityDTO.LoginResponse, error) {
	accessToken, refreshToken, userEntity, err := s.loginUseCase.Execute(ctx, idToken)
	if err != nil {
		return nil, err
	}

	userResponse := mapper.FromUserEntityToUserResponse(userEntity)

	loginResponse := &securityDTO.LoginResponse{
		User:         userResponse,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return loginResponse, nil
}

func (s *userService) ChangePassword(
    ctx context.Context, 
    userID string,
    req *securityDTO.ChangePasswordRequest,
) error {
	newPassword := req.NewPassword
	err := s.changePasswordUC.Execute(ctx, userID, newPassword)
	if err != nil {
		return err
	}

	return nil
}