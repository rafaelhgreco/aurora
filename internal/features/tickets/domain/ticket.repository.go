package domain

import (
	"context"
)

// TicketLotRepository - Focado em gerenciar os lotes de ingressos.
type TicketLotRepository interface {
	Save(ctx context.Context, lot *Ticket) (*Ticket, error)
	FindByID(ctx context.Context, id string) (*Ticket, error)
	ListByEventID(ctx context.Context, eventID string) ([]*Ticket, error)
}

// PurchasedTicketRepository - Focado nos ingressos dos usuários.
type PurchasedTicketRepository interface {
	Save(ctx context.Context, ticket *Ticket) (*Ticket, error)
	FindByID(ctx context.Context, id string) (*Ticket, error)
	ListByUserID(ctx context.Context, userID string) ([]*Ticket, error)
	ListByOrderID(ctx context.Context, orderID string) ([]*Ticket, error)
	UpdateStatus(ctx context.Context, id string, status TicketStatus) error
}
