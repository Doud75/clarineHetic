package domain

import "github.com/google/uuid"

type Instrument struct {
    UUID uuid.UUID `json:"uuid"`
    Name string    `json:"name"`
}
