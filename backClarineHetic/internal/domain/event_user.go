package domain

import "github.com/google/uuid"

type EventUser struct {
    UUID    uuid.UUID `json:"uuid"`
    UserID  uuid.UUID `json:"user_id"`
    EventID uuid.UUID `json:"event_id"`
}
