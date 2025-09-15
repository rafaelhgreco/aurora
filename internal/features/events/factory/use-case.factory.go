package factory

import (
	"aurora.com/aurora-backend/internal/features/events/domain"
	usecase "aurora.com/aurora-backend/internal/features/events/use-case"
)

type UseCaseFactory struct {
	CreateEvent *usecase.CreateEventUseCase
}

func NewUseCaseFactory(
	eventRepo domain.EventRepository,
	) *UseCaseFactory {
	return &UseCaseFactory{
		CreateEvent: usecase.NewCreateEventUseCase(eventRepo),
	}
}