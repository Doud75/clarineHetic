package usecase

import (
    "backClarineHetic/internal/domain"
)

type EventUseCase interface {
    GetEvent(searchTerm string) ([]*domain.Event, error)
}

type eventUseCase struct {
    eventRepo domain.EventRepository
}

func NewEventUseCase(eventRepo domain.EventRepository) EventUseCase {
    return &eventUseCase{
        eventRepo: eventRepo,
    }
}

func (e *eventUseCase) GetEvent(searchTerm string) ([]*domain.Event, error) {
    if searchTerm == "" {
        return e.eventRepo.GetEvent()
    }
    return e.eventRepo.GetEventWithTerm(searchTerm)
}
