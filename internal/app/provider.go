package app

import (
	httpInterface "github.com/kurniajigunawan/mikrotik-portal/internal/presenter"
	"github.com/kurniajigunawan/mikrotik-portal/internal/presenter/handler"
	middleware "github.com/kurniajigunawan/mikrotik-portal/internal/presenter/middleware"

	bridge "github.com/kurniajigunawan/mikrotik-portal/internal/bridge"

	eventRepo "github.com/kurniajigunawan/mikrotik-portal/internal/repository/event"
	homePageRepo "github.com/kurniajigunawan/mikrotik-portal/internal/repository/homepage"
	serviceRepo "github.com/kurniajigunawan/mikrotik-portal/internal/repository/service"

	eventUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/event"
	mikrotikUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/mikrotik"
	renderUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/render"

	"github.com/google/wire"
)

var (
	driverSet = wire.NewSet(
		bridge.NewRouterOSClient,
		bridge.NewDatabaseClient,
	)

	repositorySet = wire.NewSet(
		homePageRepo.New,
		eventRepo.New,
		serviceRepo.New,
	)

	usecaseSet = wire.NewSet(
		mikrotikUC.New,
		renderUC.New,
		eventUC.New,
	)

	httpSet = wire.NewSet(
		driverSet,
		repositorySet,
		usecaseSet,

		middleware.NewMiddleware,
		handler.NewHandler,
		httpInterface.NewHTTPService,
	)
)
