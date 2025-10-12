package dto

type CreateOrderRequest struct {
	UserId      string  `json:"user_id" binding:"required"`
	EventId     string  `json:"event_id" binding:"required"`
	TotalAmount float64 `json:"total_amount" binding:"required"`
}
