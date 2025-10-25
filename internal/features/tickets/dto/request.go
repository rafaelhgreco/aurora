package dto

// @Description Representa o request de compra de um ticket
type PurchaseTicketRequest struct {
	// @Id do evento
	EventId       string  `json:"event_id" binding:"required"`
	// @Id do usuário
	UserId        string  `json:"user_id" binding:"required"`
	// @Id da compra
	OrderId       string  `json:"order_id" binding:"required"`
	// @100.00
	PurchasePrice float64 `json:"purchase_price" binding:"required"`
	// @base64
	QRCodeData    string  `json:"qrcode_data"`
	// @2023-01-01T00:00:00Z
	ValidUntil    string  `json:"valid_until"`
	// @1
	Quantity      int     `json:"quantity" binding:"required,min=1"`
}
