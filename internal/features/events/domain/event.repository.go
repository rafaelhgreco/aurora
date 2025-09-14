package domain

import (
	"context"
)

type EventRepository interface {
    Save(ctx context.Context, event *Event) (*Event, error)
    FindByID(ctx context.Context, id string) (*Event, error)
    Update(ctx context.Context, event *Event) (*Event, error)
    Delete(ctx context.Context, id string) error
    ListAll(ctx context.Context) ([]*Event, error)
    FindByTitle(ctx context.Context, title string) ([]*Event, error)
}