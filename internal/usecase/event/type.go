package event

import (
	eventRepo "github.com/kurniajigunawan/mikrotik-portal/internal/repository/event"
)

type AddEventListenerRequest struct {
	ServiceID int64
	EventType int8
	Value     interface{}
}

type ListenEventsResponse struct {
	Events []ListenEvent
}

func (l *ListenEventsResponse) FromEvents(events []eventRepo.Event) {
	for _, e := range events {
		l.Events = append(l.Events, ListenEvent{
			ID:        e.ID,
			EventType: e.EventType,
			Value:     e.Value,
		})
	}
}

type ListenEvent struct {
	ID        int64
	EventType int8
	Value     interface{}
}

type EventCallbackRequest struct {
	EventID   int64
	ServiceID int64
	Status    int8
}
