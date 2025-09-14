package gateway

import (
	"context"
	"fmt"
	"log"

	"aurora.com/aurora-backend/internal/features/events/domain"
	"aurora.com/aurora-backend/internal/firebase"
	"cloud.google.com/go/firestore"
)

const ticketLotCollection = "ticket_lots"
const purchasedTicketCollection = "purchased_tickets"

type TicketLotFirestoreRepository struct {
	client *firestore.Client
}

type PurchasedTicketFirestoreRepository struct {
	client *firestore.Client
}

// NewTicketLotFirestoreRepository cria uma nova instância do repositório do Firestore.
func NewTicketLotFirestoreRepository(fbApp *firebase.FirebaseApp) (domain.TicketLotRepository, error) {
	client, err := fbApp.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Falha ao criar cliente do Firestore: %v", err)
		return nil, err
	}
	return &TicketLotFirestoreRepository{client: client}, nil
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

// Métodos para TicketLotFirestoreRepository
func (r *TicketLotFirestoreRepository) Save(ctx context.Context, lot *domain.TicketLot) (*domain.TicketLot, error) {
	_, err := r.client.Collection(ticketLotCollection).Doc(lot.ID).Set(ctx, lot)
	if err != nil {
		return nil, err
	}
	return lot, nil
}

func (r *TicketLotFirestoreRepository) FindByID(ctx context.Context, id string) (*domain.TicketLot, error) {
	docSnap, err := r.client.Collection(ticketLotCollection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var lot domain.TicketLot
	if err := docSnap.DataTo(&lot); err != nil {
		return nil, err
	}
	lot.ID = docSnap.Ref.ID
	return &lot, nil
}

func (r *TicketLotFirestoreRepository) ListByEventID(ctx context.Context, eventID string) ([]*domain.TicketLot, error) {
	iter := r.client.Collection(ticketLotCollection).Where("EventID", "==", eventID).Documents(ctx)
	defer iter.Stop()

	var lots []*domain.TicketLot
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var lot domain.TicketLot
		if err := doc.DataTo(&lot); err != nil {
			return nil, err
		}
		lot.ID = doc.Ref.ID
		lots = append(lots, &lot)
	}
	return lots, nil
}

func (r *TicketLotFirestoreRepository) DecrementAvailableQuantity(ctx context.Context, id string, amount int) error {
	docRef := r.client.Collection(ticketLotCollection).Doc(id)

	return r.client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(docRef)
		if err != nil {
			return err
		}

		var lot domain.TicketLot
		if err := doc.DataTo(&lot); err != nil {
			return err
		}

		if lot.AvailableQuantity < amount {
			return fmt.Errorf("insufficient ticket quantity: available %d, requested %d", lot.AvailableQuantity, amount)
		}

		lot.AvailableQuantity -= amount

		return tx.Set(docRef, lot)
	})
}

// Métodos para PurchasedTicketFirestoreRepository
func (r *PurchasedTicketFirestoreRepository) Save(ctx context.Context, ticket *domain.PurchasedTicket) (*domain.PurchasedTicket, error) {
	_, err := r.client.Collection(purchasedTicketCollection).Doc(ticket.ID).Set(ctx, ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (r *PurchasedTicketFirestoreRepository) FindByID(ctx context.Context, id string) (*domain.PurchasedTicket, error) {
	docSnap, err := r.client.Collection(purchasedTicketCollection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var ticket domain.PurchasedTicket
	if err := docSnap.DataTo(&ticket); err != nil {
		return nil, err
	}
	ticket.ID = docSnap.Ref.ID
	return &ticket, nil
}

func (r *PurchasedTicketFirestoreRepository) ListByUserID(ctx context.Context, userID string) ([]*domain.PurchasedTicket, error) {
	iter := r.client.Collection(purchasedTicketCollection).Where("UserID", "==", userID).Documents(ctx)
	defer iter.Stop()

	var tickets []*domain.PurchasedTicket
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var ticket domain.PurchasedTicket
		if err := doc.DataTo(&ticket); err != nil {
			return nil, err
		}
		ticket.ID = doc.Ref.ID
		tickets = append(tickets, &ticket)
	}
	return tickets, nil
}

func (r *PurchasedTicketFirestoreRepository) ListByOrderID(ctx context.Context, orderID string) ([]*domain.PurchasedTicket, error) {
	iter := r.client.Collection(purchasedTicketCollection).Where("OrderID", "==", orderID).Documents(ctx)
	defer iter.Stop()

	var tickets []*domain.PurchasedTicket
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var ticket domain.PurchasedTicket
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