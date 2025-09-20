package domain

import "context"

// OrderRepository - Focado nos pedidos/transações.
type OrderRepository interface {
    Save(ctx context.Context, order *Order) (*Order, error)
    FindByID(ctx context.Context, id string) (*Order, error)
    ListByUserID(ctx context.Context, userID string) ([]*Order, error)
}	