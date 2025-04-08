package domain

import (
    "github.com/google/uuid"
)

type Conversation struct {
    UUID    uuid.UUID `json:"uuid"`
    UserIDA uuid.UUID `json:"user_id_a"`
    UserIDB uuid.UUID `json:"user_id_b"`
}
