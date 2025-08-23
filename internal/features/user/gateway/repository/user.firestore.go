package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"aurora.com/aurora-backend/internal/features/user/domain"
	"aurora.com/aurora-backend/internal/features/user/repository"
	"aurora.com/aurora-backend/internal/firebase"
)

const userCollection = "users" // Nome da nossa coleção no Firestore

// UserFirestoreRepository é a implementação que usa o Firestore.
type UserFirestoreRepository struct {
	client *firestore.Client
}

// NewUserFirestoreRepository cria uma nova instância do repositório do Firestore.
func NewUserFirestoreRepository(fbApp *firebase.FirebaseApp) (repository.UserRepository, error) {
	client, err := fbApp.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Falha ao criar cliente do Firestore: %v", err)
		return nil, err
	}
	return &UserFirestoreRepository{client: client}, nil
}

// Save cria um novo documento de usuário no Firestore.
func (r *UserFirestoreRepository) Save(ctx context.Context, user *domain.User) (*domain.User, error) {
	_, err := r.client.Collection(userCollection).Doc(user.ID).Set(ctx, user)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

// FindByID busca um documento de usuário pelo seu ID.
func (r *UserFirestoreRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	docSnap, err := r.client.Collection(userCollection).Doc(id).Get(ctx)
	if err != nil {
		// Converte o erro para um status gRPC para verificar se é "Não Encontrado"
		if status.Code(err) == codes.NotFound {
			return nil, &domain.ErrUserNotFound{ID: id} // Um erro de domínio mais específico
		}
		return nil, err
	}

	var user domain.User
	if err := docSnap.DataTo(&user); err != nil {
		return nil, err
	}
	// O ID não está nos dados do documento, então preenchemos manualmente
	user.ID = docSnap.Ref.ID
	return &user, nil
}

// Update atualiza um documento de usuário existente no Firestore.
func (r *UserFirestoreRepository) Update(ctx context.Context, user *domain.User) (*
domain.User, error) {
	_, err := r.client.Collection(userCollection).Doc(user.ID).Set(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Delete remove um documento de usuário pelo seu ID.
func (r *UserFirestoreRepository) Delete(ctx context.Context, id string) error {
	_, err := r.client.Collection(userCollection).Doc(id).Delete(ctx)
	return err
}