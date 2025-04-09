package domain

import (
    "github.com/google/uuid"
    "time"
)

type Message struct {
    UUID           uuid.UUID `json:"uuid"`
    Content        string    `json:"content"`
    InsertAt       time.Time `json:"insert_at"`
    UserID         uuid.UUID `json:"user_id"`
    ConversationId uuid.UUID `json:"conversation_id"`
}

type MessageRepository interface {
    Create(message *Message) error
    GetMessagesByConversationID(conversationID uuid.UUID) ([]*Message, error)
}
