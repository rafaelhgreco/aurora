package usecase

import (
	"context"
	"time"

	"aurora.com/aurora-backend/internal/features/events/domain"
)

type SoftDeleteEventUseCase struct {
	repo domain.EventRepository
}
func NewSoftDeleteEventUseCase(repo domain.EventRepository) *SoftDeleteEventUseCase {
	return &SoftDeleteEventUseCase{repo: repo}
}

func (uc *SoftDeleteEventUseCase) Execute(ctx context.Context, id string) (*domain.Event, error) {
    event, err := uc.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    event.Status = domain.EVENT_CANCELLED
    event.UpdatedAt = time.Now()
    err = uc.repo.SoftDelete(ctx, id)
    if err != nil {
        return nil, err
    }
    return event, nil
}