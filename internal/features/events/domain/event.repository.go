package domain

import (
	"context"
)

type EventRepository interface {
    Save(ctx context.Context, event *Event) (*Event, error)
    FindByID(ctx context.Context, id string) (*Event, error)
    Update(ctx context.Context, event *Event) (*Event, error)
    SoftDelete(ctx context.Context, id string) error
    ListAll(ctx context.Context, filter map[string]interface{}) ([]*Event, error)
    FindByTitle(ctx context.Context, title string) ([]*Event, error)
}