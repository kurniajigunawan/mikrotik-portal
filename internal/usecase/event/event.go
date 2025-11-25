package event

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	gerr "github.com/aidapedia/gdk/error"
	ghttp "github.com/aidapedia/gdk/http"

	"github.com/bytedance/sonic"
	"github.com/kurniajigunawan/mikrotik-portal/internal/repository/event"
)

// AddEventListener is a function to add event listener
func (u *Usecase) AddEventListener(ctx context.Context, req AddEventListenerRequest) error {
	var (
		err error
	)
	valByte, err := sonic.Marshal(req.Value)
	if err != nil {
		return gerr.NewWithMetadata(err, ghttp.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
	}
	err = u.eventRepo.CreateEvent(ctx, event.CreateEventRequest{
		ServiceID: req.ServiceID,
		EventType: req.EventType,
		Value:     json.RawMessage(valByte),
		Status:    event.StatusWaitingToConsume,
	})
	if err != nil {
		return gerr.NewWithMetadata(err, ghttp.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
	}
	return nil
}

func (u *Usecase) ListenEvents(ctx context.Context, serviceID int64) (ListenEventsResponse, error) {
	events, err := u.eventRepo.GetActiveEventsByServiceID(ctx, serviceID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ListenEventsResponse{}, nil
		}
		return ListenEventsResponse{}, gerr.NewWithMetadata(err, ghttp.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
	}

	resp := ListenEventsResponse{}
	resp.FromEvents(events)
	return resp, nil
}

// EventCallback is a function to callback event
func (u *Usecase) EventCallback(ctx context.Context, req EventCallbackRequest) error {
	evtSts, err := event.CheckStatus(req.Status)
	if err != nil {
		return gerr.NewWithMetadata(err, ghttp.Metadata(http.StatusBadRequest, "Invalid Status"))
	}
	err = u.eventRepo.SetStatus(ctx, req.EventID, req.ServiceID, evtSts)
	if err != nil {
		return gerr.NewWithMetadata(err, ghttp.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
	}
	return nil
}
