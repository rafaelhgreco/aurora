package dto

type PurchaseTicketRequest struct {
	EventId       string  `json:"event_id" binding:"required"`
	UserId        string  `json:"user_id" binding:"required"`
	OrderId       string  `json:"order_id" binding:"required"`
	PurchasePrice float64 `json:"purchase_price" binding:"required"`
	QRCodeData    string  `json:"qrcode_data"`
	ValidUntil    string  `json:"valid_until"`
	Quantity      int     `json:"quantity" binding:"required,min=1"`
}
