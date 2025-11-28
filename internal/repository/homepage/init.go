package homepage

import (
	"context"
	"database/sql"
)

type Interface interface {
	GetActiveMenu(ctx context.Context) ([]Menu, error)
}

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) Interface {
	return &Repository{
		db: db,
	}
}
