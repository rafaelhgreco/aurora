package factory

import (
	eventsDomain "aurora.com/aurora-backend/internal/features/events/domain"
	"aurora.com/aurora-backend/internal/features/order/domain"
	usecase "aurora.com/aurora-backend/internal/features/order/use-case"
)

type UseCaseFactory struct {
	CreateOrder *usecase.CreateOrderUseCase
}

func NewUseCaseFactory(orderRepo domain.OrderRepository, eventRepo eventsDomain.EventRepository) *UseCaseFactory {
	return &UseCaseFactory{
		CreateOrder: usecase.NewCreateOrderUseCase(orderRepo, eventRepo),
	}
}
