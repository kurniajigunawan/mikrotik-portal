package main

import (
	"context"

	"github.com/aidapedia/gdk/log"
	"github.com/kurniajigunawan/mikrotik-portal/internal/app"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	ctx := context.Background()
	log.New(&log.Config{
		Level:       log.LoggerLevel("info"),
		DefaultTags: map[string]interface{}{},
	})
	defer log.Sync()
	service := app.InitHTTPServer(ctx)
	errs := service.ListenAndServe()
	if errs != nil {
		log.ErrorCtx(ctx, "Failed to start HTTP server", zap.Error(errs))
	}
}
