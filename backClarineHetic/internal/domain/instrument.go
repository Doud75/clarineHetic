package domain

import "github.com/google/uuid"

type instrument struct {
    UUID uuid.UUID `json:"uuid"`
    Name string    `json:"name"`
}
