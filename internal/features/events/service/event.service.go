package service

import (
	"context"

	"aurora.com/aurora-backend/internal/features/events/dto"
	"aurora.com/aurora-backend/internal/features/events/mapper"

	usecase "aurora.com/aurora-backend/internal/features/events/use-case/event"
)

type eventService struct {
	createEventUseCase  *usecase.CreateEventUseCase
	FindByIDUseCase	*usecase.FindByIDEventUseCase
	ListAllEventUseCase *usecase.ListAllEventUsecase
	SoftDeleteEventUseCase *usecase.SoftDeleteEventUseCase
}

func NewEventService(
	createEventUC *usecase.CreateEventUseCase,
	findByIDUC *usecase.FindByIDEventUseCase,
	listAllUC *usecase.ListAllEventUsecase,
	deleteUc *usecase.SoftDeleteEventUseCase,
) IEventService {
	return &eventService{
		createEventUseCase:    createEventUC,
		FindByIDUseCase:	findByIDUC,
		ListAllEventUseCase:   listAllUC,
		SoftDeleteEventUseCase: deleteUc,
	}
}

func (s *eventService) Save(ctx context.Context, req *dto.CreateEventRequest) error {
    eventEntity, err := mapper.FromCreateRequestToEventEntity(req)
    if err != nil {
        return err
    }
    _, err = s.createEventUseCase.Execute(ctx, eventEntity)
    if err != nil {
        return err
    }
    return nil
}

func (s *eventService) FindByID(ctx context.Context ,id string) (req *dto.EventResponse, err error) {
	eventEntity, err := s.FindByIDUseCase.Execute(ctx, id)
	if err != nil {
		return nil, err
	}
	eventDTO, err := mapper.FromEventEntityToResponse(eventEntity)
	if err != nil {
		return nil, err
	}
	return eventDTO, nil
}

func (s *eventService) ListEvents(ctx context.Context, filter map[string]interface{}) ([]*dto.EventResponse, error) {
	events, err := s.ListAllEventUseCase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	eventDTOs := make([]*dto.EventResponse, len(events))
	for i, event := range events {
		eventDTO, err := mapper.FromEventEntityToResponse(event)
		if err != nil {
			return nil, err
		}
		eventDTOs[i] = eventDTO
	}
	return eventDTOs, nil
}

func (s *eventService) SoftDeleteEvent(ctx context.Context, id string) error {
    return s.SoftDeleteEventUseCase.Execute(ctx, id)
}