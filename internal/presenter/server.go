package presenter

import (
	"fmt"

	"github.com/aidapedia/gdk/http/server"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/kurniajigunawan/mikrotik-portal/internal/presenter/handler"
)

// HTTPServiceInterface is an interface to handle http service
type HTTPServiceInterface interface {
	ListenAndServe() error
}

// HTTPService is a struct to handle http service
type HTTPService struct {
	svr *server.Server
}

// NewHTTPService is a function to create a new http service
func NewHTTPService(handler *handler.Handler) HTTPServiceInterface {
	engine := html.New("./public", ".html")
	svr, _ := server.NewWithDefaultConfig("mikrotik-portal", server.WithAppConfig(fiber.Config{
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
		StrictRouting: true,
		CaseSensitive: true,
		Immutable:     true,
		Views:         engine,
	}))
	svr.Get("/:page", handler.Render)

	api := svr.Group("/api", cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	api.Post("/reset-session", handler.ResetSession)

	return &HTTPService{
		svr: svr,
	}
}

// ListenAndServe is a function to start http service
func (h *HTTPService) ListenAndServe() error {
	return h.svr.Listen(fmt.Sprintf("%s:%d", "0.0.0.0", 8080))
}
