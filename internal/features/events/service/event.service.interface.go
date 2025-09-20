package service

import (
	"context"

	"aurora.com/aurora-backend/internal/features/events/dto"
)

type IEventService interface {
	Save(ctx context.Context, req *dto.CreateEventRequest ) (error)
	FindByID(ctx context.Context, id string) (req *dto.EventResponse, err error)
	ListEvents(ctx context.Context, filter map[string]interface{}) ([]*dto.EventResponse, error)
	SoftDeleteEvent(ctx context.Context, id string) error
	// UpdateEvent(id string, data interface{}) error
}