package usecase

import (
	"context"
	"time"

	"aurora.com/aurora-backend/internal/features/events/domain"
)

type ListAllEventUsecase struct {
	repo domain.EventRepository
}

func NewListAllEventUsecase(repo domain.EventRepository) *ListAllEventUsecase {
	return &ListAllEventUsecase{
		repo: repo,
	}
}
func (uc *ListAllEventUsecase) Execute(ctx context.Context, filter map[string]interface{}) ([]*domain.Event, error) {
	events, err := uc.repo.ListAll(ctx, filter)
	for _, event := range events {
		if status := event.DetermineStatus(); status != event.Status && event.Status != domain.EVENT_CANCELLED {
			event.Status = status
			event.UpdatedAt = time.Now()
			_, err = uc.repo.Update(ctx, event)
		}
	}
	if err != nil {
		return nil, err
	}

	return events, nil
}
