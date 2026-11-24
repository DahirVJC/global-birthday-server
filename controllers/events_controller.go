package controllers

import (
	"fmt"
	"global-birthday-server/models"
	"global-birthday-server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventsController struct {
	eventsService   *services.EventsService
	sessionsService *services.SessionsService
}

func NewEventsController(eventsService *services.EventsService, sessionsService *services.SessionsService) *EventsController {
	return &EventsController{
		eventsService:   eventsService,
		sessionsService: sessionsService,
	}
}

func (ec *EventsController) userIdFromToken(c *gin.Context) (uuid.UUID, bool) {
	accessToken := c.GetHeader("access-token")

	userId, err := ec.sessionsService.GetUser(accessToken)
	if err != nil {
		parseAndRespond(c, err)
		return uuid.Nil, false
	}

	return userId, true
}

func parseEventIdParam(c *gin.Context) (uuid.UUID, bool) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		parseAndRespond(c, fmt.Errorf("malformed id"))
		return uuid.Nil, false
	}

	return id, true
}

func (ec *EventsController) GetEvents(c *gin.Context) {
	userId, ok := ec.userIdFromToken(c)
	if !ok {
		return
	}

	includePast, err := strconv.ParseBool(c.DefaultQuery("past", "false"))
	if err != nil {
		parseAndRespond(c, fmt.Errorf("malformed past"))
		return
	}

	events, err := ec.eventsService.List(userId, includePast)
	if err != nil {
		parseAndRespond(c, err)
		return
	}

	c.JSON(http.StatusOK, events)
}

func (ec *EventsController) GetEvent(c *gin.Context) {
	userId, ok := ec.userIdFromToken(c)
	if !ok {
		return
	}

	id, ok := parseEventIdParam(c)
	if !ok {
		return
	}

	event, err := ec.eventsService.Get(userId, id)
	if err != nil {
		parseAndRespond(c, err)
		return
	}

	c.JSON(http.StatusOK, event)
}

func (ec *EventsController) PostEvent(c *gin.Context) {
	userId, ok := ec.userIdFromToken(c)
	if !ok {
		return
	}

	var newEvent models.EventReq
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		parseAndRespond(c, err)
		return
	}

	if err := ec.eventsService.Create(userId, newEvent); err != nil {
		parseAndRespond(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (ec *EventsController) PutEvent(c *gin.Context) {
	userId, ok := ec.userIdFromToken(c)
	if !ok {
		return
	}

	id, ok := parseEventIdParam(c)
	if !ok {
		return
	}

	var newEventData models.EventReq
	if err := c.ShouldBindJSON(&newEventData); err != nil {
		parseAndRespond(c, err)
		return
	}

	if err := ec.eventsService.Update(userId, id, newEventData); err != nil {
		parseAndRespond(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (ec *EventsController) DeleteEvent(c *gin.Context) {
	userId, ok := ec.userIdFromToken(c)
	if !ok {
		return
	}

	id, ok := parseEventIdParam(c)
	if !ok {
		return
	}

	if err := ec.eventsService.Delete(userId, id); err != nil {
		parseAndRespond(c, err)
		return
	}

	c.Status(http.StatusOK)
}
