package controller

import (
	"net/http"

	"aurora.com/aurora-backend/internal/features/tickets/dto"
	"aurora.com/aurora-backend/internal/features/tickets/mapper"
	usecase "aurora.com/aurora-backend/internal/features/tickets/use-case"
	"github.com/gin-gonic/gin"
)

type TicketController struct {
	createTicket *usecase.PurchaseTicketUseCase
}

func NewTicketController(createTicket *usecase.PurchaseTicketUseCase) *TicketController {
	return &TicketController{createTicket: createTicket}
}

func (ctrl *TicketController) CreateTicket(c *gin.Context) {
	var req dto.PurchaseTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticketDomain, err := mapper.FromPurchaseTicketRequestToDomain(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = ctrl.createTicket.Execute(c.Request.Context(), ticketDomain)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
