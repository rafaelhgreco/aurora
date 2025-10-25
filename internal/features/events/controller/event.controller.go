package controller

import (
	"net/http"

	"aurora.com/aurora-backend/internal/features/events/dto"
	"aurora.com/aurora-backend/internal/features/events/mapper"
	usecase "aurora.com/aurora-backend/internal/features/events/use-case/event"
	sharedErrors "aurora.com/aurora-backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	createEvent     *usecase.CreateEventUseCase
	findByIDEvent   *usecase.FindByIDEventUseCase
	listAllEvent    *usecase.ListAllEventUsecase
	softDeleteEvent *usecase.SoftDeleteEventUseCase
}

func NewEventController(createEvent *usecase.CreateEventUseCase, findByIDEvent *usecase.FindByIDEventUseCase, listAllEvent *usecase.ListAllEventUsecase, softDeleteEvent *usecase.SoftDeleteEventUseCase) *EventController {
	return &EventController{
		createEvent:     createEvent,
		findByIDEvent:   findByIDEvent,
		listAllEvent:    listAllEvent,
		softDeleteEvent: softDeleteEvent,
	}
}

// CreateEvent godoc
// @Summary Criar novo evento
// @Description Cria um novo evento no sistema
// @Tags events
// @Accept json
// @Produce json
// @Param request body dto.CreateEventRequest true "Dados do evento"
// @Success 201 "Event created successfully"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /events [post]
func (ctrl *EventController) CreateEvent(c *gin.Context) {
	var req dto.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sharedErrors.HandleError(c, sharedErrors.ErrInvalidInput)
		return
	}
	eventEntity, err := mapper.FromCreateRequestToEventEntity(&req)
	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}
	_, err = ctrl.createEvent.Execute(c.Request.Context(), eventEntity)

	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}
	c.Status(http.StatusCreated)
}

// GetEvent godoc
// @Summary Buscar evento por ID
// @Description Retorna os dados de um evento específico
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "ID do evento"
// @Success 200 {object} dto.EventResponse
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Event not found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /events/{id} [get]
func (ctrl *EventController) GetEvent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		sharedErrors.HandleError(c, sharedErrors.ErrInvalidInput)
		return
	}
	eventEntity, err := ctrl.findByIDEvent.Execute(c.Request.Context(), id)
	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}
	eventResponse, err := mapper.FromEventEntityToResponse(eventEntity)
	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, eventResponse)
}

// ListEvents godoc
// @Summary Listar todos os eventos
// @Description Retorna uma lista de todos os eventos disponíveis
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {array} dto.EventResponse
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /events [get]
func (ctrl *EventController) ListEvents(c *gin.Context) {
	filter := make(map[string]interface{})
	events, err := ctrl.listAllEvent.Execute(c.Request.Context(), filter)
	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}
	eventsResponse, err := mapper.FromEventEntitiesToResponses(events)
	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, eventsResponse)
}

// SoftDeleteEvent godoc
// @Summary Deletar evento (soft delete)
// @Description Remove logicamente um evento do sistema
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "ID do evento"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /events/{id} [delete]
func (ctrl *EventController) SoftDeleteEvent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		sharedErrors.HandleError(c, sharedErrors.ErrInvalidInput)
		return
	}
	eventEntity, err := ctrl.softDeleteEvent.Execute(c.Request.Context(), id)
	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}
	_, err = mapper.FromSoftDeleteEventEntity(eventEntity)
	if err != nil {
		sharedErrors.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
