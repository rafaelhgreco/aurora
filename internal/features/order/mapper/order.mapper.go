package mapper

import (
	"errors"

	"aurora.com/aurora-backend/internal/features/order/domain"
	"aurora.com/aurora-backend/internal/features/order/dto"
)

func FromCreateOrderRequestToDomain(req *dto.CreateOrderRequest) (*domain.Order, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}
	return &domain.Order{
		UserId:      req.UserId,
		EventId:     req.EventId,
		TotalAmount: req.TotalAmount,
	}, nil
}
