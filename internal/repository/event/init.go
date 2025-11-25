package event

import (
	"context"
	"database/sql"
)

type Interface interface {
	CreateEvent(ctx context.Context, event CreateEventRequest) error
	SetStatus(ctx context.Context, eventID, serviceID int64, status Status) error
	GetActiveEventsByServiceID(ctx context.Context, serviceID int64) ([]Event, error)
}

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) Interface {
	return &Repository{
		db: db,
	}
}
