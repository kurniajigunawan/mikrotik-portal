package handler

import (
	"errors"
	"net/http"

	gerr "github.com/aidapedia/gdk/error"
	ghttp "github.com/aidapedia/gdk/http"
	"github.com/aidapedia/gdk/util"
	"github.com/gofiber/fiber/v3"
	"github.com/kurniajigunawan/mikrotik-portal/internal/presenter/handler/model"
)

func (h *Handler) AddEventListener(c fiber.Ctx) error {
	var (
		ctx = c.Context()
		req model.AddEventListenerRequest
	)
	if err := c.Bind().Body(&req); err != nil {
		ghttp.JSONResponse(c, nil, gerr.NewWithMetadata(err, ghttp.Metadata(http.StatusBadRequest, "Bad Request")))
		return err
	}

	err := h.eventUsecase.AddEventListener(ctx, req.ToUsecaseRequest())
	if err != nil {
		ghttp.JSONResponse(c, nil, err)
		return err
	}

	ghttp.JSONResponse(c, nil, nil)
	return nil
}

func (h *Handler) ListenEvents(c fiber.Ctx) error {
	var (
		ctx  = c.Context()
		resp = model.ListenEventsResponse{}
	)
	serviceID := c.Locals("ServiceID")
	if serviceID == "" {
		ghttp.JSONResponse(c, nil, gerr.NewWithMetadata(errors.New("ServiceID is required"), ghttp.Metadata(http.StatusBadRequest, "ServiceID is required")))
		return nil
	}
	event, err := h.eventUsecase.ListenEvents(ctx, util.ToInt64(serviceID))
	if err != nil {
		ghttp.JSONResponse(c, nil, err)
		return err
	}

	resp.FromEvents(event)
	ghttp.JSONResponse(c, &ghttp.SuccessResponse{
		Data: resp.Events,
	}, nil)
	return nil
}

func (h *Handler) CallbackEvent(c fiber.Ctx) error {
	var (
		ctx = c.Context()
		req model.EventCallbackRequest
	)
	if err := c.Bind().Body(&req); err != nil {
		ghttp.JSONResponse(c, nil, gerr.NewWithMetadata(err, ghttp.Metadata(http.StatusBadRequest, "Bad Request")))
		return err
	}

	serviceID := c.Locals("ServiceID")
	if serviceID == "" {
		ghttp.JSONResponse(c, nil, gerr.NewWithMetadata(errors.New("ServiceID is required"), ghttp.Metadata(http.StatusBadRequest, "ServiceID is required")))
		return nil
	}

	err := h.eventUsecase.EventCallback(ctx, req.ToUsecaseRequest(util.ToInt64(serviceID)))
	if err != nil {
		ghttp.JSONResponse(c, nil, err)
		return err
	}

	ghttp.JSONResponse(c, nil, nil)
	return nil
}
