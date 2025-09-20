package factory

import (
	"aurora.com/aurora-backend/internal/features/events/domain"
	usecase "aurora.com/aurora-backend/internal/features/events/use-case/event"
)

type UseCaseFactory struct {
	CreateEvent *usecase.CreateEventUseCase
	FindByIDEvent *usecase.FindByIDEventUseCase
	ListAllEvent *usecase.ListAllEventUsecase
}

func NewUseCaseFactory(
	eventRepo domain.EventRepository,
	) *UseCaseFactory {
	return &UseCaseFactory{
		CreateEvent: usecase.NewCreateEventUseCase(eventRepo),
		FindByIDEvent: usecase.NewFindByIDEventUseCase(eventRepo),
		ListAllEvent: usecase.NewListAllEventUsecase(eventRepo),
	}
}