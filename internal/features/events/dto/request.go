package dto

type CreateEventRequest struct {
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description" binding:"required"`
	StartTime    string `json:"start_time" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	EndTime      string `json:"end_time" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Location     string `json:"location" binding:"required"`
	TotalTickets int    `json:"total_tickets" binding:"required,min=1"`
}

type UpdateEventRequest struct {
	Title        *string `json:"title"`
	Description  *string `json:"description"`
	Date         *string `json:"date" binding:"omitempty,datetime=2006-01-02"`
	Location     *string `json:"location"`
	TotalTickets *int    `json:"total_tickets" binding:"omitempty,min=1"`
}

type GetEventByIDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}
