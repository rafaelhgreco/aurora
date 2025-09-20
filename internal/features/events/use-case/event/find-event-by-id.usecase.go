package usecase

import (
	"context"

	"aurora.com/aurora-backend/internal/features/events/domain"
)

type FindByIDEventUseCase struct {
	repo domain.EventRepository
}

func NewFindByIDEventUseCase(repo domain.EventRepository) *FindByIDEventUseCase {
	return &FindByIDEventUseCase{repo: repo}
}

func (uc *FindByIDEventUseCase) Execute(ctx context.Context, id string) (*domain.Event, error) {
	event, err := uc.repo.FindByID(ctx, id)
	if status := event.DetermineStatus(); status != event.Status {
		event.Status = status
		_, err = uc.repo.Update(ctx, event)
	}
	if err != nil {
		return nil, err
	}
	return event, nil
}