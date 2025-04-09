package dto

import (
    "backClarineHetic/internal/domain"
    "github.com/google/uuid"
)

type MessageRequest struct {
    Content string `json:"content" binding:"required"`
}

type ConversationResponse struct {
    ConversationID uuid.UUID         `json:"conversation_id"`
    Messages       []*domain.Message `json:"messages"`
}
