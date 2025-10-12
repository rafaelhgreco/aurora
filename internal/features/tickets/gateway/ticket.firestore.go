package gateway

import (
	"context"
	"log"

	"aurora.com/aurora-backend/internal/features/tickets/domain"
	"aurora.com/aurora-backend/internal/firebase"
	"cloud.google.com/go/firestore"
)

const ticketLotCollection = "ticket_lots"
const purchasedTicketCollection = "purchased_tickets"

type PurchasedTicketFirestoreRepository struct {
	client *firestore.Client
}

// NewPurchasedTicketFirestoreRepository cria uma nova instância do repositório do Firestore.
func NewPurchasedTicketFirestoreRepository(fbApp *firebase.FirebaseApp) (domain.PurchasedTicketRepository, error) {
	client, err := fbApp.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Falha ao criar cliente do Firestore: %v", err)
		return nil, err
	}
	return &PurchasedTicketFirestoreRepository{client: client}, nil
}

func (r *PurchasedTicketFirestoreRepository) Save(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	_, err := r.client.Collection(purchasedTicketCollection).Doc(ticket.ID).Set(ctx, ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (r *PurchasedTicketFirestoreRepository) FindByID(ctx context.Context, id string) (*domain.Ticket, error) {
	docSnap, err := r.client.Collection(purchasedTicketCollection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var ticket domain.Ticket
	if err := docSnap.DataTo(&ticket); err != nil {
		return nil, err
	}
	ticket.ID = docSnap.Ref.ID
	return &ticket, nil
}

func (r *PurchasedTicketFirestoreRepository) ListByUserID(ctx context.Context, userID string) ([]*domain.Ticket, error) {
	iter := r.client.Collection(purchasedTicketCollection).Where("UserID", "==", userID).Documents(ctx)
	defer iter.Stop()

	var tickets []*domain.Ticket
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var ticket domain.Ticket
		if err := doc.DataTo(&ticket); err != nil {
			return nil, err
		}
		ticket.ID = doc.Ref.ID
		tickets = append(tickets, &ticket)
	}
	return tickets, nil
}

func (r *PurchasedTicketFirestoreRepository) ListByOrderID(ctx context.Context, orderID string) ([]*domain.Ticket, error) {
	iter := r.client.Collection(purchasedTicketCollection).Where("OrderID", "==", orderID).Documents(ctx)
	defer iter.Stop()

	var tickets []*domain.Ticket
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var ticket domain.Ticket
		if err := doc.DataTo(&ticket); err != nil {
			return nil, err
		}
		ticket.ID = doc.Ref.ID
		tickets = append(tickets, &ticket)
	}
	return tickets, nil
}

func (r *PurchasedTicketFirestoreRepository) UpdateStatus(ctx context.Context, id string, status domain.TicketStatus) error {
	_, err := r.client.Collection(purchasedTicketCollection).Doc(id).Update(ctx, []firestore.Update{
		{Path: "Status", Value: status},
	})
	return err
}
