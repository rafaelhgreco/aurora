package controller

import (
	"net/http"

	"aurora.com/aurora-backend/internal/features/events/dto"
	service "aurora.com/aurora-backend/internal/features/events/service"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	eventService service.IEventService
}

func NewEventController(eventService service.IEventService) *EventController {
	return &EventController{
		eventService: eventService,
	}
}

func (ctrl *EventController) CreateEvent(c *gin.Context) {
	var req dto.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.eventService.Save(c.Request.Context(), &req)
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
	eventResponse, err := ctrl.eventService.FindByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventResponse)
}

func (ctrl *EventController) ListEvents(c *gin.Context) {
	filter := make(map[string]interface{})
	events, err := ctrl.eventService.ListEvents(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}