package mapper

import (
	"errors"
	"time"

	"aurora.com/aurora-backend/internal/features/events/domain"
	"aurora.com/aurora-backend/internal/features/events/dto"
)

func FromCreateRequestToEventEntity(req *dto.CreateEventRequest) (*domain.Event, error) {
	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		return nil, errors.New("invalid start_time format, must be RFC3339")
	}
	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		return nil, errors.New("invalid end_time format, must be RFC3339")
	}
	if endTime.Before(startTime) {
		return nil, errors.New("end_time must be after start_time")
	}
	return &domain.Event{
		Title:            req.Title,
		Description:      req.Description,
		StartTime:        startTime,
		EndTime:          endTime,
		Location:         req.Location,
		TotalTickets:     req.TotalTickets,
		AvailableTickets: req.TotalTickets,
		Status:           domain.EVENT_SCHEDULED,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}, nil
}

func FromEventEntityToResponse(entity *domain.Event) (*dto.EventResponse, error) {
	if entity == nil {
		return nil, errors.New("event entity is nil")
	}
	return &dto.EventResponse{
		ID:               entity.ID,
		Title:            entity.Title,
		Description:      entity.Description,
		Location:         entity.Location,
		TotalTickets:     entity.TotalTickets,
		AvailableTickets: entity.AvailableTickets,
		StartTime:        entity.StartTime,
		EndTime:          entity.EndTime,
		Status:           entity.Status.String(),
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
	}, nil
}

func FromEventEntitiesToResponses(entities []*domain.Event) ([]*dto.EventResponse, error) {
	responses := make([]*dto.EventResponse, 0, len(entities))
	for _, entity := range entities {
		resp, err := FromEventEntityToResponse(entity)
		if err != nil {
			return nil, err
		}
		responses = append(responses, resp)
	}
	return responses, nil
}

func FromSoftDeleteEventEntity(entity *domain.Event) (*dto.EventResponse, error) {
	if entity == nil {
		return nil, errors.New("event entity is nil")
	}
	return &dto.EventResponse{
		ID:               entity.ID,
		Title:            entity.Title,
		Description:      entity.Description,
		Location:         entity.Location,
		TotalTickets:     entity.TotalTickets,
		AvailableTickets: entity.AvailableTickets,
		StartTime:        entity.StartTime,
		EndTime:          entity.EndTime,
		Status:           entity.Status.String(),
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
	}, nil
}
