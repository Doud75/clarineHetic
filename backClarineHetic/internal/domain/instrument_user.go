package domain

import "github.com/google/uuid"

type InstrumentUser struct {
    UUID         uuid.UUID `json:"uuid"`
    UserID       uuid.UUID `json:"user_id"`
    InstrumentID uuid.UUID `json:"instrument_id"`
}
