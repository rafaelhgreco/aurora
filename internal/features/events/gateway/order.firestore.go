package gateway

import (
	"context"
	"log"

	"aurora.com/aurora-backend/internal/features/events/domain"
	"aurora.com/aurora-backend/internal/firebase"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const orderCollection = "orders"


type OrderFirestoreRepository struct {
	client *firestore.Client
}

// NewOrderFirestoreRepository cria uma nova instância do repositório do Firestore.
func NewOrderFirestoreRepository(fbApp *firebase.FirebaseApp) (domain.OrderRepository, error) {
	client, err := fbApp.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Falha ao criar cliente do Firestore: %v", err)
		return nil, err
	}
	return &OrderFirestoreRepository{client: client}, nil
}

func (r *OrderFirestoreRepository) Save(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	_, err := r.client.Collection(orderCollection).Doc(order.ID).Set(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderFirestoreRepository) FindByID(ctx context.Context, id string) (*domain.Order, error) {
	docSnap, err := r.client.Collection(orderCollection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var order domain.Order
	if err := docSnap.DataTo(&order); err != nil {
		return nil, err
	}
	order.ID = docSnap.Ref.ID
	return &order, nil
}

func (r *OrderFirestoreRepository) ListByUserID(ctx context.Context, userID string) ([]*domain.Order, error) {
	iter := r.client.Collection(orderCollection).Where("user_id", "==", userID).Documents(ctx)
	var orders []*domain.Order
	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, err
		}
		var order domain.Order
		if err := doc.DataTo(&order); err != nil {
			return nil, err
		}
		order.ID = doc.Ref.ID
		orders = append(orders, &order)
	}
	return orders, nil
}