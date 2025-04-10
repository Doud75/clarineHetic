package controller

import (
    "backClarineHetic/internal/usecase"
    "github.com/gin-gonic/gin"
    "net/http"
)

type EventController interface {
    GetEvent(ctx *gin.Context)
}

type eventController struct {
    eventUseCase usecase.EventUseCase
}

func NewEventController(eventUseCase usecase.EventUseCase) EventController {
    return &eventController{eventUseCase: eventUseCase}
}

func (e *eventController) GetEvent(c *gin.Context) {
    searchTerm := c.Query("search_term")

    event, err := e.eventUseCase.GetEvent(searchTerm)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": event})
}
