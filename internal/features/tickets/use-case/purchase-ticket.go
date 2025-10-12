package usecase

import (
	"context"
	"errors"
	"time"

	eventsDomain "aurora.com/aurora-backend/internal/features/events/domain"
	orderDomain "aurora.com/aurora-backend/internal/features/order/domain"
	"aurora.com/aurora-backend/internal/features/tickets/domain"
	"github.com/google/uuid"
)

type PurchaseTicketUseCase struct {
	eventRepo  eventsDomain.EventRepository
	orderRepo  orderDomain.OrderRepository
	ticketRepo domain.PurchasedTicketRepository
}

func NewPurchaseTicketUseCase(
	eventRepo eventsDomain.EventRepository,
	orderRepo orderDomain.OrderRepository,
	ticketRepo domain.PurchasedTicketRepository,
) *PurchaseTicketUseCase {
	return &PurchaseTicketUseCase{
		eventRepo:  eventRepo,
		orderRepo:  orderRepo,
		ticketRepo: ticketRepo,
	}
}

func (uc *PurchaseTicketUseCase) Execute(ctx context.Context, req *domain.Ticket) (*domain.Ticket, error) {
	event, err := uc.eventRepo.FindByID(ctx, req.EventId)
	if err != nil {
		return nil, errors.New("event not found")
	}
	if event.AvailableTickets <= 0 {
		return nil, errors.New("event is sold out")
	}

	order, err := uc.orderRepo.FindByID(ctx, req.OrderId)
	if err != nil {
		return nil, errors.New("order not found")
	}
	if order.EventId != req.EventId || order.UserId != req.UserId {
		return nil, errors.New("order not found")
	}
	event.AvailableTickets -= 1
	event.UpdatedAt = time.Now()
	_, err = uc.eventRepo.Update(ctx, event)
	if err != nil {
		return nil, errors.New("failed to update event")
	}

	validUntil := req.ValidUntil

	ticket := &domain.Ticket{
		ID:            uuid.New().String(),
		OrderId:       req.OrderId,
		EventId:       req.EventId,
		UserId:        req.UserId,
		PurchasePrice: req.PurchasePrice,
		PurchaseDate:  time.Now(),
		QRCodeData:    req.QRCodeData,
		Status:        domain.TicketValid,
		IssuedAt:      time.Now(),
		ValidUntil:    validUntil,
	}
	_, err = uc.ticketRepo.Save(ctx, ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil

}
