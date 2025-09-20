package service

import (
	"context"

	"aurora.com/aurora-backend/internal/features/events/dto"
)

type IEventService interface {
	Save(ctx context.Context, req *dto.CreateEventRequest ) (error)
	// FindByID(id string) (interface{}, error)
	// UpdateEvent(id string, data interface{}) error
	// DeleteEvent(id string) error
	// ListEvents(filter map[string]interface{}) ([]interface{}, error)
	// FindByTitle(title string) (interface{}, error)
}