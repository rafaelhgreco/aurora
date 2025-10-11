package domain

import (
	"time"
)

type TicketStatus int

const (
    TicketValid   TicketStatus = iota 
    TicketUsed                    
    TicketInvalid 
	ticketCanceled               
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