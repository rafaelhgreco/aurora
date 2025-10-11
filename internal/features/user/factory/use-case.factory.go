package factory

import (
	"aurora.com/aurora-backend/internal/features/user/domain"
	securityUseCase "aurora.com/aurora-backend/internal/features/user/security/use-case"
	usecase "aurora.com/aurora-backend/internal/features/user/use-case"
)

type UseCaseFactory struct {
	CreateUser     *usecase.CreateUserUseCase
	UpdateUser     *usecase.UpdateUserUseCase
	GetUserByID    *usecase.GetUserByIDUseCase
	DeleteUser     *usecase.DeleteUserUseCase
	LoginUser      *securityUseCase.LoginUserUseCase
	ChangePassword *securityUseCase.ChangePasswordUseCase
}

func NewUseCaseFactory(
	userRepo domain.UserRepository,
	authClient domain.AuthClient,
) *UseCaseFactory {
	return &UseCaseFactory{
		CreateUser:  usecase.NewCreateUserUseCase(userRepo, authClient),
		GetUserByID: usecase.NewGetUserByIDUseCase(userRepo),
		UpdateUser:  usecase.NewUpdateUserUseCase(userRepo),
		DeleteUser:  usecase.NewDeleteUserUseCase(userRepo),

		LoginUser:      securityUseCase.NewLoginUserUseCase(userRepo, authClient),
		ChangePassword: securityUseCase.NewChangePasswordUseCase(authClient),
	}
}
