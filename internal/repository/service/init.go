package service

import (
	"context"
	"database/sql"
)

type Interface interface {
	GetServiceByClientID(ctx context.Context, clientID string) (Service, error)
	GetActiveServices(ctx context.Context) ([]Service, error)
}

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) Interface {
	return &Repository{
		db: db,
	}
}
