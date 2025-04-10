package controller

import (
    "backClarineHetic/internal/adapter/dto"
    "backClarineHetic/internal/domain"
    "backClarineHetic/internal/usecase"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "time"
)

type EventController interface {
    GetEvent(ctx *gin.Context)
    CreateEvent(ctx *gin.Context)
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

func (e *eventController) CreateEvent(c *gin.Context) {
    var req dto.EventRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    authUserUUIDStr, exists := c.Get("userUUID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non authentifié"})
        return
    }
    authUserUUID, err := uuid.Parse(authUserUUIDStr.(string))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Le uuid n'est pas un UUID valide"})
        return
    }

    layout := "2006-01-02 15:04:05"
    startDate, err := time.Parse(layout, req.StartDate)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Format de start_date invalide"})
        return
    }

    event := domain.Event{
        Name:      req.Name,
        Adress:    req.Adress,
        City:      req.City,
        StartDate: startDate,
        UserID:    authUserUUID,
    }

    if err = e.eventUseCase.CreateEvent(&event); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": event})

}
