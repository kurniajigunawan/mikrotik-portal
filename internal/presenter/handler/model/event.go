package model

import "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/event"

type AddEventListenerRequest struct {
	ServiceID int64       `json:"service_id"`
	EventType int8        `json:"event_type"`
	Value     interface{} `json:"value"`
}

func (a *AddEventListenerRequest) ToUsecaseRequest() event.AddEventListenerRequest {
	return event.AddEventListenerRequest{
		ServiceID: a.ServiceID,
		EventType: a.EventType,
		Value:     a.Value,
	}
}

type ListenEventsResponse struct {
	Events []ListenEvent `json:"event"`
}

type ListenEvent struct {
	ID        int64       `json:"id"`
	EventType int8        `json:"event_type"`
	Value     interface{} `json:"value"`
}

func (l *ListenEventsResponse) FromEvents(events event.ListenEventsResponse) {
	l.Events = []ListenEvent{}
	for _, e := range events.Events {
		l.Events = append(l.Events, ListenEvent{
			ID:        e.ID,
			EventType: e.EventType,
			Value:     e.Value,
		})
	}
}

type EventCallbackRequest struct {
	EventID int64 `json:"event_id"`
	Status  int8  `json:"status"`
}

func (e *EventCallbackRequest) ToUsecaseRequest(serviceID int64) event.EventCallbackRequest {
	return event.EventCallbackRequest{
		EventID:   e.EventID,
		Status:    e.Status,
		ServiceID: serviceID,
	}
}
