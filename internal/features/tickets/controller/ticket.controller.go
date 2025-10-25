package controller

import (
	"net/http"

	"aurora.com/aurora-backend/internal/features/tickets/dto"
	"aurora.com/aurora-backend/internal/features/tickets/mapper"
	usecase "aurora.com/aurora-backend/internal/features/tickets/use-case"
	sharedErrors "aurora.com/aurora-backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type TicketController struct {
	createTicket *usecase.PurchaseTicketUseCase
}

func NewTicketController(createTicket *usecase.PurchaseTicketUseCase) *TicketController {
	return &TicketController{createTicket: createTicket}
}
// CreateTicket godoc
// @Summary Criar ingressos
// @Description Cria N ingressos para um evento com base em uma ordem
// @Tags tickets
// @Accept json
// @Produce json
// @Param request body dto.PurchaseTicketRequest true "Dados de compra de ingressos"
// @Success 201 {array} dto.TicketPurchasedResponse "Lista de ingressos criados"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 422 {object} map[string]string "Unprocessable Entity"
// @Router /tickets [post]
func (ctrl *TicketController) CreateTicket(c *gin.Context) {
	var req dto.PurchaseTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sharedErrors.HandleError(c, sharedErrors.ErrInvalidInput)
		return
	}

	ticketDomain, err := mapper.FromPurchaseTicketRequestToDomain(&req)
	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}
	tickets, err := ctrl.createTicket.Execute(c.Request.Context(), ticketDomain)
	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}
	ticketResponses := mapper.FromDomainTicketsToResponses(tickets)
	c.JSON(http.StatusCreated, ticketResponses)
}
