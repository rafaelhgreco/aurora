package usecase

import (
	"context"

	"aurora.com/aurora-backend/internal/features/events/domain"
	"github.com/google/uuid"
)

type CreateEventUseCase struct {
	repo domain.EventRepository
}

func NewCreateEventUseCase(repo domain.EventRepository) *CreateEventUseCase {
	return &CreateEventUseCase{repo: repo}
}

func (uc *CreateEventUseCase) Execute(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	event.ID = uuid.New().String()

	savedEvent, err := uc.repo.Save(ctx, event)
	if err != nil {
		return nil, err
	}
	return savedEvent, nil
}
