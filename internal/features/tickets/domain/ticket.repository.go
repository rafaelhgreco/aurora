package domain

import (
	"context"
)

type PurchasedTicketRepository interface {
	Save(ctx context.Context, ticket *Ticket) (*Ticket, error)
	FindByID(ctx context.Context, id string) (*Ticket, error)
	ListByUserID(ctx context.Context, userID string) ([]*Ticket, error)
	ListByOrderID(ctx context.Context, orderID string) ([]*Ticket, error)
	UpdateStatus(ctx context.Context, id string, status TicketStatus) error
}
