package bridge

import (
	"context"
	"os"

	"github.com/aidapedia/go-routeros"
)

var Hook []func() error

func NewRouterOSClient(ctx context.Context) *routeros.RouterOS {
	routerBuilder := routeros.NewRouterOS(&routeros.Options{
		Address:       os.Getenv("ROUTEROS_ADDRESS"),
		Username:      os.Getenv("ROUTEROS_USERNAME"),
		Password:      os.Getenv("ROUTEROS_PASSWORD"),
		AutoReconnect: true,
	})
	// err := routerBuilder.Connect()
	// if err != nil {
	// 	// log.FatalCtx(ctx, gerror.New(err).Error())
	// }
	Hook = append(Hook, routerBuilder.Close)
	return routerBuilder
}
