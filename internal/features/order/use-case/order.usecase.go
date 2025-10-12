package usecase

import (
	"context"
	"errors"
	"time"

	eventsDomain "aurora.com/aurora-backend/internal/features/events/domain"
	"aurora.com/aurora-backend/internal/features/order/domain"
	"github.com/google/uuid"
)

type CreateOrderUseCase struct {
	orderRepo domain.OrderRepository
	eventRepo eventsDomain.EventRepository
}

func NewCreateOrderUseCase(orderRepo domain.OrderRepository, eventRepo eventsDomain.EventRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepo: orderRepo,
		eventRepo: eventRepo,
	}
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, req *domain.Order) (string, error) {
	event, err := uc.eventRepo.FindByID(ctx, req.EventId)
	if err != nil {
		return "", errors.New("evento não encontrado")
	}
	if event.AvailableTickets <= 0 {
		return "", errors.New("ingressos esgotados")
	}

	order := &domain.Order{
		ID:          uuid.New().String(),
		UserId:      req.UserId,
		EventId:     req.EventId,
		OrderDate:   time.Now(),
		TotalAmount: req.TotalAmount,
		Status:      domain.ORDER_PENDING,
	}

	_, err = uc.orderRepo.Save(ctx, order)
	if err != nil {
		return "", errors.New("falha ao criar order")
	}

	return order.ID, nil
}
