package controller

import (
    "backClarineHetic/internal/adapter/dto"
    "backClarineHetic/internal/domain"
    "backClarineHetic/internal/usecase"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "time"
)

type ConversationController interface {
    GetConversation(ctx *gin.Context)
    SaveMessage(ctx *gin.Context)
}

type conversationController struct {
    conversationUseCase usecase.ConversationUseCase
}

func NewConversationController(conversationUseCase usecase.ConversationUseCase) ConversationController {
    return &conversationController{conversationUseCase: conversationUseCase}
}

func (cc *conversationController) GetConversation(c *gin.Context) {
    userUuidStr := c.Query("user_uuid")
    if userUuidStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Le paramètre user_uuid est requis"})
        return
    }
    userUuid, err := uuid.Parse(userUuidStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Le paramètre user_uuid n'est pas un UUID valide"})
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

    convResponse, err := cc.conversationUseCase.GetOrCreateConv(userUuid, authUserUUID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": convResponse})
}

func (cc *conversationController) SaveMessage(c *gin.Context) {
    var req dto.MessageRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide", "details": err.Error()})
        return
    }

    uuidParamStr := c.Param("uuid")
    uuidParam, err := uuid.Parse(uuidParamStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Le paramètre uuid n'est pas un UUID valide"})
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

    message := domain.Message{
        Content:        req.Content,
        InsertAt:       time.Now(),
        UserID:         authUserUUID,
        ConversationId: uuidParam,
    }

    fmt.Println(req)

    if err = cc.conversationUseCase.SaveMessage(&message); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Message saved successfully",
        "data":    message,
    })
}
