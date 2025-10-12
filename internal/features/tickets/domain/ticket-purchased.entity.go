package domain

import (
	"time"
)

type TicketStatus int

const (
	TicketValid TicketStatus = iota
	TicketUsed
	TicketInvalid
	ticketCanceled
)

type Ticket struct {
	ID            string       `firestore:"-"`
	OrderId       string       `firestore:"orderId"`
	EventId       string       `firestore:"eventId"`
	UserId        string       `firestore:"userId"`
	PurchasePrice float64      `firestore:"purchasePrice"`
	PurchaseDate  time.Time    `firestore:"purchaseDate"`
	QRCodeData    string       `firestore:"qrCodeData"`
	Status        TicketStatus `firestore:"status"`
	Quantity      int          `firestore:"quantity"`
	IssuedAt      time.Time    `firestore:"issuedAt"`
	ValidUntil    time.Time    `firestore:"validUntil"`
}
