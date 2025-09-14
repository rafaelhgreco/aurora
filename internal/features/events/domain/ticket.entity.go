package domain

import (
	"time"
)

type TicketLot struct {
    ID               string    `firestore:"-"`
    EventID          string    `firestore:"eventId"`
    Name             string    `firestore:"name"`           
    Price            float64   `firestore:"price"`
    TotalQuantity    int       `firestore:"totalQuantity"`   
    AvailableQuantity int      `firestore:"availableQuantity"`
    CreatedAt        time.Time `firestore:"createdAt"`
    UpdatedAt        time.Time `firestore:"updatedAt"`
}

type TicketStatus int

const (
    TICKET_VALID   TicketStatus = iota 
    TICKET_USED                     
    TICKET_INVALID                 
)

type PurchasedTicket struct {
    ID            string       `firestore:"-"`
    OrderID       string       `firestore:"orderId"`       
    EventID       string       `firestore:"eventId"`
    UserID        string       `firestore:"userId"`
    TicketLotID   string       `firestore:"ticketLotId"`   
    PurchasePrice float64      `firestore:"purchasePrice"` 
    PurchaseDate  time.Time    `firestore:"purchaseDate"`
    QRCodeData    string       `firestore:"qrCodeData"`    
    Status        TicketStatus `firestore:"status"`
    IssuedAt      time.Time    `firestore:"issuedAt"`      
    ValidUntil    time.Time    `firestore:"validUntil"`   
}