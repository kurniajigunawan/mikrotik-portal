package event

import (
	"context"

	eventRepo "github.com/kurniajigunawan/mikrotik-portal/internal/repository/event"
)

type Interface interface {
	AddEventListener(ctx context.Context, req AddEventListenerRequest) error
	ListenEvents(ctx context.Context, serviceID int64) (ListenEventsResponse, error)
	EventCallback(ctx context.Context, req EventCallbackRequest) error
}

type Usecase struct {
	eventRepo eventRepo.Interface
}

func New(eventRepo eventRepo.Interface) Interface {
	return &Usecase{
		eventRepo: eventRepo,
	}
}
