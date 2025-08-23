package usecase

import (
	"context"

	"aurora.com/aurora-backend/internal/features/user/domain"
)

type CreateUserUseCase struct {
	userRepo   domain.UserRepository
	authClient domain.AuthClient
	// O hasher também pode ser removido se o Firebase já está cuidando da senha.
}

func NewCreateUserUseCase(repo domain.UserRepository, authClient domain.AuthClient) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: repo, authClient: authClient}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, user *domain.User) (*domain.User, error) {
	// 1. Delega a criação do usuário no Auth diretamente para o gateway,
	//    passando a nossa entidade de domínio.
	uid, err := uc.authClient.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// 2. O resto da lógica permanece o mesmo.
	user.ID = uid
	user.Password = "" // Limpa a senha antes de salvar no nosso DB

	savedUser, err := uc.userRepo.Save(ctx, user)
	if err != nil {
		// Lógica de rollback (deletar usuário do Auth) poderia ser adicionada aqui
		return nil, err
	}

	return savedUser, nil
}