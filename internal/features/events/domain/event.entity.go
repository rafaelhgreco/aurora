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

func (es EventStatus) String() string {
    switch es {
    case EVENT_SCHEDULED:
        return "SCHEDULED"
    case EVENT_OPEN_FOR_SALE:
        return "OPEN_FOR_SALE"
    case EVENT_SOLD_OUT:
        return "SOLD_OUT"
    case EVENT_CANCELLED:
        return "CANCELLED"
    case EVENT_FINISHED:
        return "FINISHED"
    default:
        return "UNKNOWN"
    }
}

type Event struct {
	ID          string `firestore:"-"`
	Title       string `firestore:"title"`
	Description string `firestore:"description"`
	StartTime   time.Time `firestore:"start_time"`
	EndTime     time.Time `firestore:"end_time"`
	Location    string `firestore:"location"`
	TotalTickets      int         `firestore:"totalTickets"`    
    AvailableTickets  int         `firestore:"availableTickets"`          
	CreatedAt   time.Time `firestore:"created_at"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}

func (e *Event) DetermineStatus() EventStatus {
	now := time.Now()
	if e.AvailableTickets <= 0 {
		return EVENT_SOLD_OUT
	}
	if now.Before(e.StartTime) {
		return EVENT_SCHEDULED
	}
	if now.After(e.EndTime) {
		return EVENT_FINISHED
	}
	return EVENT_OPEN_FOR_SALE
}