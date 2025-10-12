package mapper

import (
	"time"

	"aurora.com/aurora-backend/internal/features/tickets/domain"
	"aurora.com/aurora-backend/internal/features/tickets/dto"
)

func FromPurchaseTicketRequestToDomain(req *dto.PurchaseTicketRequest) (*domain.Ticket, error) {
	validUntil, err := time.Parse(time.RFC3339, req.ValidUntil)
	if err != nil {
		return nil, err
	}
	return &domain.Ticket{
		EventId:       req.EventId,
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		PurchasePrice: req.PurchasePrice,
		QRCodeData:    req.QRCodeData,
		ValidUntil:    validUntil,
	}, nil
}
