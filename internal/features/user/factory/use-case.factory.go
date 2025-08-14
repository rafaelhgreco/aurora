package factory

import (
	"aurora.com/aurora-backend/internal/features/user/domain"
	"aurora.com/aurora-backend/internal/features/user/repository"
	usecase "aurora.com/aurora-backend/internal/features/user/use-case"
)

type UseCaseFactory struct {
	CreateUser *usecase.CreateUserUseCase
	GetUser    *usecase.GetUserByIDUseCase
}

func NewUseCaseFactory(userRepo repository.UserRepository, passwordHasher domain.PasswordHasher) *UseCaseFactory {
	return &UseCaseFactory{
		CreateUser: usecase.NewCreateUserUseCase(userRepo, passwordHasher),
		GetUser:    usecase.NewGetUserByIDUseCase(userRepo),
	}
}