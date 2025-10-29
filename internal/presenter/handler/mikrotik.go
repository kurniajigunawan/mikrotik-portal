package handler

import (
	"net/http"

	gers "github.com/aidapedia/gdk/error"
	ghttp "github.com/aidapedia/gdk/http"
	"github.com/gofiber/fiber/v3"
	"github.com/kurniajigunawan/mikrotik-portal/internal/presenter/handler/model"
	pkgLog "github.com/kurniajigunawan/mikrotik-portal/pkg/log"
)

func (h *Handler) ResetSession(c fiber.Ctx) error {
	var (
		ctx = c.Context()
		req model.ResetSessionRequest
	)
	if err := c.Bind().Body(&req); err != nil {
		ghttp.JSONResponse(c, nil, gers.NewWithMetadata(err, pkgLog.Metadata(http.StatusBadRequest, "Bad Request")))
		return err
	}

	err := h.mikrotikUsecase.ResetSession(ctx, req.Username)
	if err != nil {
		ghttp.JSONResponse(c, nil, err)
		return err
	}

	ghttp.JSONResponse(c, nil, nil)
	return nil
}
