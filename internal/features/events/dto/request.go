package dto

// @Description Representa o request de criação de um novo evento
type CreateEventRequest struct {
	// @Titulo do evento
	Title        string `json:"title" binding:"required"`
	// @Descrição do evento
	Description  string `json:"description" binding:"required"`
	// @@2023-01-01T00:00:00Z
	StartTime    string `json:"start_time" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	// @@2023-01-01T00:00:00Z
	EndTime      string `json:"end_time" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	// @Local do evento
	Location     string `json:"location" binding:"required"`
	// @300
	TotalTickets int    `json:"total_tickets" binding:"required,min=1"`
}

// @Description Representa o request de atualização de um evento
type UpdateEventRequest struct {
	// @Titulo do evento
	Title        *string `json:"title"`
	// @Descrição do evento
	Description  *string `json:"description"`
	// @2023-01-01T00:00:00Z
	Date         *string `json:"date" binding:"omitempty,datetime=2006-01-02"`
	// @Local do evento
	Location     *string `json:"location"`
	// @300
	TotalTickets *int    `json:"total_tickets" binding:"omitempty,min=1"`
}

// @Description Representa o request de busca de um evento por ID
type GetEventByIDRequest struct {
	// @Id do evento
	ID string `uri:"id" binding:"required,uuid"`
}
