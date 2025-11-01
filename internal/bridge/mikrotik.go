package bridge

import (
	"log"
	"os"

	"github.com/aidapedia/go-routeros"
)

var Hook []func() error

func NewRouterOSClient() *routeros.RouterOS {
	routerBuilder := routeros.NewRouterOS(&routeros.Options{
		Address:       os.Getenv("ROUTEROS_ADDRESS"),
		Username:      os.Getenv("ROUTEROS_USERNAME"),
		Password:      os.Getenv("ROUTEROS_PASSWORD"),
		AutoReconnect: true,
	})
	err := routerBuilder.Connect()
	if err != nil {
		log.Fatal(err)
	}
	Hook = append(Hook, routerBuilder.Close)
	return routerBuilder
}
