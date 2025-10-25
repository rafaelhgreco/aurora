package dto

// @Description Representa o request de criação de uma nova ordem
type CreateOrderRequest struct {
	// @ID do usuário
	UserId      string  `json:"user_id" binding:"required"`
	// @ID do evento
	EventId     string  `json:"event_id" binding:"required"`
	// @100.00
	TotalAmount float64 `json:"total_amount" binding:"required"`
	// @1
	Quantity    int     `json:"quantity" binding:"required,min=1"`
}
