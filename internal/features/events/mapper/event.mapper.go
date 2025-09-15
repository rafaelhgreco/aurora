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
	return &domain.Event{
		Title:        req.Title,
		Description:  req.Description,
		StartTime:    startTime,
		EndTime:      endTime,
		Location:     req.Location,
		TotalTickets: req.TotalTickets,
	}, nil
}

func FromEventEntityToEventResponse(entity *domain.Event) *dto.EventResponse {
	return &dto.EventResponse{
		ID:              entity.ID,
		Title:           entity.Title,
		Description:     entity.Description,
		Location:        entity.Location,
		TotalTickets:    entity.TotalTickets,
		AvailableTickets: entity.AvailableTickets,
		Status:          entity.Status.String(),
		StartTime: 	 entity.StartTime,
		EndTime: 	 entity.EndTime,
		CreatedAt:       entity.CreatedAt,
		UpdatedAt:       entity.UpdatedAt,
	}

}