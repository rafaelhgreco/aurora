package usecase

import (
	"context"

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
	if err != nil {
		return nil, err
	}
	return events, nil
}