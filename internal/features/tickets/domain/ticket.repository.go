package domain

import (
	"context"
)

// TicketLotRepository - Focado em gerenciar os lotes de ingressos.
type TicketLotRepository interface {
    Save(ctx context.Context, lot *TicketLot) (*TicketLot, error)
    FindByID(ctx context.Context, id string) (*TicketLot, error)
    ListByEventID(ctx context.Context, eventID string) ([]*TicketLot, error)
    DecrementAvailableQuantity(ctx context.Context, id string, amount int) error
}

// PurchasedTicketRepository - Focado nos ingressos dos usuários.
type PurchasedTicketRepository interface {
    Save(ctx context.Context, ticket *PurchasedTicket) (*PurchasedTicket, error)
    FindByID(ctx context.Context, id string) (*PurchasedTicket, error)
    ListByUserID(ctx context.Context, userID string) ([]*PurchasedTicket, error)
    ListByOrderID(ctx context.Context, orderID string) ([]*PurchasedTicket, error)
    UpdateStatus(ctx context.Context, id string, status TicketStatus) error
}