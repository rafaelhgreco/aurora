package factory

import (
	eventsDomain "aurora.com/aurora-backend/internal/features/events/domain"
	orderDomain "aurora.com/aurora-backend/internal/features/order/domain"
	"aurora.com/aurora-backend/internal/features/tickets/domain"
	usecase "aurora.com/aurora-backend/internal/features/tickets/use-case"
)

type UseCaseFactory struct {
	PurchaseTicket *usecase.PurchaseTicketUseCase
}

func NewUseCaseFactory(
	eventRepo eventsDomain.EventRepository,
	orderRepo orderDomain.OrderRepository,
	ticketRepo domain.PurchasedTicketRepository) *UseCaseFactory {
	return &UseCaseFactory{
		PurchaseTicket: usecase.NewPurchaseTicketUseCase(eventRepo, orderRepo, ticketRepo),
	}
}
