package dto

import (
	"time"

	ticketStatus "aurora.com/aurora-backend/internal/features/tickets/domain"
)

type TicketResponse struct {
	ID         string                    `json:"id"`
	Title      string                    `json:"title"`
	Content    string                    `json:"content"`
	CreatedAt  time.Time                 `json:"created_at"`
	UpdatedAt  time.Time                 `json:"updated_at"`
	DeletedAt  *time.Time                `json:"deleted_at,omitempty"`
	QRCodeData string                    `json:"qrcode_data"`
	ValidUntil time.Time                 `json:"valid_until"`
	Status     ticketStatus.TicketStatus `json:"status"`
}

type TicketPurchasedResponse struct {
	QRCodeData    string       `json:"qrcode_data"`
	Status        ticketStatus.TicketStatus `json:"status"`
	Quantity      int          `json:"quantity"`
	IssuedAt      time.Time    `json:"issued_at"`
	ValidUntil    time.Time    `json:"valid_until"`
}