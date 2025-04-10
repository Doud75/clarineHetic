package dto

type EventRequest struct {
    Name      string `json:"name" binding:"required"`
    Adress    string `json:"adress" binding:"required"`
    City      string `json:"city" binding:"required"`
    StartDate string `json:"start_date" binding:"required"`
}
