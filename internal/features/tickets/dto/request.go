package dto

type PurchaseTicketRequest struct {
	OrderID       string  `json:"order_id" binding:"required"`
	EventID       string  `json:"event_id" binding:"required"`
	UserID        string  `json:"user_id" binding:"required"`
	TicketLotID   string  `json:"ticket_lot_id" binding:"required"`
	PurchasePrice float64 `json:"purchase_price" binding:"required"`
}
