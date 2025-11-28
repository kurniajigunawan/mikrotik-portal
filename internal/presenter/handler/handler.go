package handler

import (
	eventUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/event"
	mikrotikUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/mikrotik"
	renderUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/render"
)

type Handler struct {
	mikrotikUsecase mikrotikUC.UsecaseItf
	renderUsecase   renderUC.UsecaseItf
	eventUsecase    eventUC.Interface
}

// NewHandler is a function to create a new handler
func NewHandler(eventUsecase eventUC.Interface, mikrotikUsecase mikrotikUC.UsecaseItf, renderUsecase renderUC.UsecaseItf) *Handler {
	return &Handler{
		mikrotikUsecase: mikrotikUsecase,
		renderUsecase:   renderUsecase,
		eventUsecase:    eventUsecase,
	}
}
