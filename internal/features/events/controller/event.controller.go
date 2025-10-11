package controller

import (
	"net/http"

	"aurora.com/aurora-backend/internal/features/events/dto"
	"aurora.com/aurora-backend/internal/features/events/mapper"
	usecase "aurora.com/aurora-backend/internal/features/events/use-case/event"
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

func (ctrl *EventController) CreateEvent(c *gin.Context) {
	var req dto.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	eventEntity, err := mapper.FromCreateRequestToEventEntity(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = ctrl.createEvent.Execute(c.Request.Context(), eventEntity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (ctrl *EventController) GetEvent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	eventEntity, err := ctrl.findByIDEvent.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	eventResponse, err := mapper.FromEventEntityToResponse(eventEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventResponse)
}

func (ctrl *EventController) ListEvents(c *gin.Context) {
	filter := make(map[string]interface{})
	events, err := ctrl.listAllEvent.Execute(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	eventsResponse, err := mapper.FromEventEntitiesToResponses(events)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventsResponse)
}

func (ctrl *EventController) SoftDeleteEvent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	eventEntity, err := ctrl.softDeleteEvent.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = mapper.FromSoftDeleteEventEntity(eventEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
