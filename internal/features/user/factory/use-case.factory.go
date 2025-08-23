package factory

import (
	"aurora.com/aurora-backend/internal/features/user/domain"
	"aurora.com/aurora-backend/internal/features/user/repository"
	securityUseCase "aurora.com/aurora-backend/internal/features/user/security/use-case"
	usecase "aurora.com/aurora-backend/internal/features/user/use-case"
)

type UseCaseFactory struct {
	CreateUser     *usecase.CreateUserUseCase
	GetUser        *usecase.GetUserByIDUseCase
	UpdateUser     *usecase.UpdateUserUseCase
	DeleteUser     *usecase.DeleteUserUseCase
	LoginUser      *securityUseCase.LoginUserUseCase
	ChangePassword *securityUseCase.ChangePasswordUseCase
}

func NewUseCaseFactory(
    userRepo repository.UserRepository, 
    passwordHasher domain.PasswordHasher, 
    authClient domain.AuthClient,
) *UseCaseFactory {
	return &UseCaseFactory{
		CreateUser: usecase.NewCreateUserUseCase(userRepo, authClient),
		GetUser:    usecase.NewGetUserByIDUseCase(userRepo),
		UpdateUser: usecase.NewUpdateUserUseCase(userRepo),
		DeleteUser: usecase.NewDeleteUserUseCase(userRepo),

		LoginUser:      securityUseCase.NewLoginUserUseCase(userRepo, authClient),
		ChangePassword: securityUseCase.NewChangePasswordUseCase(authClient),
	}
}