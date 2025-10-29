package mikrotik

import (
	"context"

	"github.com/aidapedia/go-routeros"
)

type UsecaseItf interface {
	ResetSession(ctx context.Context, username string) error
}

type Usecase struct {
	routerBuilder *routeros.RouterOS
}

func New(routerBuilder *routeros.RouterOS) UsecaseItf {
	return &Usecase{routerBuilder: routerBuilder}
}
