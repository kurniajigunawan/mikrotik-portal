package app

import (
	httpInterface "github.com/kurniajigunawan/mikrotik-portal/internal/presenter"
	"github.com/kurniajigunawan/mikrotik-portal/internal/presenter/handler"

	bridge "github.com/kurniajigunawan/mikrotik-portal/internal/bridge"
	mikrotikUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/mikrotik"
	renderUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/render"

	"github.com/google/wire"
)

var (
	driverSet = wire.NewSet(
		bridge.NewRouterOSClient,
	)

	usecaseSet = wire.NewSet(
		mikrotikUC.New,
		renderUC.New,
	)

	httpSet = wire.NewSet(
		driverSet,
		usecaseSet,
		handler.NewHandler,
		httpInterface.NewHTTPService,
	)
)
