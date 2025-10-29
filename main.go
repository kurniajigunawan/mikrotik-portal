package main

import (
	"context"

	"github.com/aidapedia/gdk/log"
	"github.com/kurniajigunawan/mikrotik-portal/internal/app"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	log.New(&log.Config{
		Level:       log.LoggerLevel("info"),
		Caller:      true,
		DefaultTags: map[string]interface{}{},
	})
	defer log.Sync()
	service := app.InitHTTPServer()
	errs := service.ListenAndServe()
	if errs != nil {
		log.ErrorCtx(context.Background(), "Failed to start HTTP server", zap.Error(errs))
	}
}
