//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"
	"github.com/kurniajigunawan/mikrotik-portal/internal/presenter"
)

func InitHTTPServer() presenter.HTTPServiceInterface {
	wire.Build(httpSet)
	return &presenter.HTTPService{}
}
