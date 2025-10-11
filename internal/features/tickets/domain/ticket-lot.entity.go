package domain

import (
	"time"
)

type TicketLot struct {
	ID                string    `firestore:"-"`
	EventID           string    `firestore:"eventId"`
	Name              string    `firestore:"name"`
	Price             float64   `firestore:"price"`
	TotalQuantity     int       `firestore:"totalQuantity"`
	AvailableQuantity int       `firestore:"availableQuantity"`
	CreatedAt         time.Time `firestore:"createdAt"`
	UpdatedAt         time.Time `firestore:"updatedAt"`
}
