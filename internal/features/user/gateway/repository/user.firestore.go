package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"aurora.com/aurora-backend/internal/features/user/domain"
	"aurora.com/aurora-backend/internal/firebase"
)

const userCollection = "users"

type UserFirestoreRepository struct {
	client *firestore.Client
}

// NewUserFirestoreRepository cria uma nova instância do repositório do Firestore.
func NewUserFirestoreRepository(fbApp *firebase.FirebaseApp) (domain.UserRepository, error) {
	client, err := fbApp.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Falha ao criar cliente do Firestore: %v", err)
		return nil, err
	}
	return &UserFirestoreRepository{client: client}, nil
}

func (r *UserFirestoreRepository) Save(ctx context.Context, user *domain.User) (*domain.User, error) {
	_, err := r.client.Collection(userCollection).Doc(user.ID).Set(ctx, user)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (r *UserFirestoreRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	docSnap, err := r.client.Collection(userCollection).Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, &domain.ErrUserNotFound{ID: id}
		}
		return nil, err
	}

	var user domain.User
	if err := docSnap.DataTo(&user); err != nil {
		return nil, err
	}
	user.ID = docSnap.Ref.ID
	return &user, nil
}

func (r *UserFirestoreRepository) Update(ctx context.Context, user *domain.User) (*
domain.User, error) {
	_, err := r.client.Collection(userCollection).Doc(user.ID).Set(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserFirestoreRepository) Delete(ctx context.Context, id string) error {
	_, err := r.client.Collection(userCollection).Doc(id).Delete(ctx)
	return err
}