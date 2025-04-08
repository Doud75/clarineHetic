package domain

import (
    "github.com/google/uuid"
    "time"
)

type Event struct {
    UUID      uuid.UUID `json:"uuid"`
    Name      string    `json:"name"`
    Longitude string    `json:"longitude"`
    Latitude  string    `json:"latitude"`
    Adress    string    `json:"adress"`
    City      string    `json:"city"`
    StartDate time.Time `json:"start_date"`
    UserID    uuid.UUID `json:"user_id"`
}
