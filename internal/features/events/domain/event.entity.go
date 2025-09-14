package domain

import (
	"time"
)

type EventStatus int

const (
	EVENT_SCHEDULED EventStatus = iota
	EVENT_OPEN_FOR_SALE
	EVENT_SOLD_OUT
	EVENT_CANCELLED
	EVENT_FINISHED
)

type Event struct {
	ID          string `firestore:"-"`
	Title       string `firestore:"title"`
	Description string `firestore:"description"`
	StartTime   time.Time `firestore:"start_time"`
	EndTime     time.Time `firestore:"end_time"`
	Location    string `firestore:"location"`
	TotalTickets      int         `firestore:"totalTickets"`    
    AvailableTickets  int         `firestore:"availableTickets"`  
    Status  	EventStatus `firestore:"status"`            
	CreatedAt   time.Time `firestore:"created_at"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}