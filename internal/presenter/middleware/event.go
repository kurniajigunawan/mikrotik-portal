package middleware

import (
	"database/sql"
	"errors"

	gerr "github.com/aidapedia/gdk/error"
	ghttp "github.com/aidapedia/gdk/http"
	"github.com/aidapedia/gdk/util"
	"github.com/gofiber/fiber/v3"

	serviceRepo "github.com/kurniajigunawan/mikrotik-portal/internal/repository/service"
)

func (m *Middleware) EventMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		var (
			err error
		)
		defer func() {
			if err != nil {
				ghttp.JSONResponse(c, nil, err)
			}
		}()
		clientID := c.Get("X-Client-Id")
		if clientID == "" {
			err = gerr.NewWithMetadata(errors.New("X-Client-Id is empty"), ghttp.Metadata(fiber.StatusUnauthorized, "Invalid Token"))
			return err
		}
		service, err := m.serviceRepo.GetServiceByClientID(c.Context(), clientID)
		if err != nil {
			if err == sql.ErrNoRows {
				err = gerr.NewWithMetadata(err, ghttp.Metadata(fiber.StatusUnauthorized, "Invalid Token"))
				return err
			}
			err = gerr.NewWithMetadata(err, ghttp.Metadata(fiber.StatusInternalServerError, "Internal Server Error"))
			return err
		}
		if service.Status != serviceRepo.StatusActive {
			err = gerr.NewWithMetadata(errors.New("Service is not active"), ghttp.Metadata(fiber.StatusUnauthorized, "Invalid Token"))
			return err
		}
		if service.ClientSecret != c.Get("X-Client-Secret") {
			err = gerr.NewWithMetadata(errors.New("invalid secret"), ghttp.Metadata(fiber.StatusUnauthorized, "Invalid Token"))
			return err
		}
		c.Locals("ServiceID", util.ToStr(service.ID))
		return c.Next()
	}
}
