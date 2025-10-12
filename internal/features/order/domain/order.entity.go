package domain

import (
	"time"
)

type OrderStatus int

const (
	ORDER_PENDING OrderStatus = iota
	ORDER_COMPLETED
	ORDER_FAILED
	ORDER_CANCELLED
)

type Order struct {
	ID          string      `firestore:"-"`
	UserId      string      `firestore:"userId"`
	EventId     string      `firestore:"eventId"`
	OrderDate   time.Time   `firestore:"orderDate"`
	TotalAmount float64     `firestore:"totalAmount"`
	Status      OrderStatus `firestore:"status"`
}
