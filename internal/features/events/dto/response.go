package dto

import "time"

type EventResponse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StartTime        time.Time  `json:"date"`
	EndTime        time.Time  `json:"end_time"`
	Location    string     `json:"location"`
	AvailableTickets int        `json:"available_tickets"`
	TotalTickets    int        `json:"total_tickets"`
	Status 	 string        `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
