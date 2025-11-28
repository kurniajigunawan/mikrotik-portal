package event

import (
	"encoding/json"
	"errors"
	"time"
)

type Status int8

const (
	StatusWaitingToConsume Status = iota
	StatusConsumed
)

func CheckStatus(status int8) (Status, error) {
	switch status {
	case 0, 1:
		return Status(status), nil
	default:
		return Status(status), errors.New("invalid status")
	}
}

type Event struct {
	ID        int64
	ServiceID int64
	EventType int8
	Value     json.RawMessage
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateEventRequest struct {
	ServiceID int64
	EventType int8
	Value     json.RawMessage
	Status    Status
}
