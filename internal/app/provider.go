package app

import (
	"os"

	"github.com/aidapedia/go-routeros"
	httpInterface "github.com/kurniajigunawan/mikrotik-portal/internal/presenter"
	"github.com/kurniajigunawan/mikrotik-portal/internal/presenter/handler"

	mikrotikUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/mikrotik"

	"github.com/google/wire"
)

var Hook []func() error

func mikrotikProvider() *routeros.RouterOS {
	routerBuilder := routeros.NewRouterOS(&routeros.Options{
		Address:       os.Getenv("ROUTEROS_ADDRESS"),
		Username:      os.Getenv("ROUTEROS_USERNAME"),
		Password:      os.Getenv("ROUTEROS_PASSWORD"),
		AutoReconnect: true,
	})
	err := routerBuilder.Connect()
	if err != nil {
		// panic(err)
	}
	Hook = append(Hook, routerBuilder.Close)
	return routerBuilder
}

var (
	driverSet = wire.NewSet(
		mikrotikProvider,
	)

	usecaseSet = wire.NewSet(
		mikrotikUC.New,
	)

	httpSet = wire.NewSet(
		driverSet,
		usecaseSet,
		handler.NewHandler,
		httpInterface.NewHTTPService,
	)
)
