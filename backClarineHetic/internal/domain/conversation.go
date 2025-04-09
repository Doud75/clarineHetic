package domain

import (
    "github.com/google/uuid"
)

type Conversation struct {
    UUID    uuid.UUID `json:"uuid"`
    UserIDA uuid.UUID `json:"user_id_a"`
    UserIDB uuid.UUID `json:"user_id_b"`
}

type ConversationRepository interface {
    Create(conversation *Conversation) error
    GetByID(conversationID uuid.UUID) (*Conversation, error)
    GetByUserIDs(authUserUUID uuid.UUID, userUUID uuid.UUID) (*Conversation, error)
    Delete(conversationID uuid.UUID) error
    Update(conversation *Conversation) error
}
