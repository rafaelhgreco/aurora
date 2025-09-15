package service

import (
	"context"

	"aurora.com/aurora-backend/internal/features/events/dto"
	"aurora.com/aurora-backend/internal/features/events/mapper"

	usecase "aurora.com/aurora-backend/internal/features/events/use-case"
)

type eventService struct {
	createEventUseCase  *usecase.CreateEventUseCase
}

func NewEventService(
	createEventUC *usecase.CreateEventUseCase,
) IEventService {
	return &eventService{
		createEventUseCase:    createEventUC,
	}
}

func (s *eventService) Save(ctx context.Context, req *dto.CreateEventRequest) (*dto.EventResponse, error) {
	eventEntity, err := mapper.FromCreateRequestToEventEntity(req)
	if err != nil {
		return nil, err
	}
	createdEvent, err := s.createEventUseCase.Execute(ctx, eventEntity)
	if err != nil {
		return nil, err
	}
	return mapper.FromEventEntityToEventResponse(createdEvent), nil
}
