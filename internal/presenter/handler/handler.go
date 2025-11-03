package handler

import (
	mikrotikUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/mikrotik"
	renderUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/render"
)

type Handler struct {
	mikrotikUsecase mikrotikUC.UsecaseItf
	renderUsecase   renderUC.UsecaseItf
}

// NewHandler is a function to create a new handler
func NewHandler(mikrotikUsecase mikrotikUC.UsecaseItf, renderUsecase renderUC.UsecaseItf) *Handler {
	return &Handler{
		mikrotikUsecase: mikrotikUsecase,
		renderUsecase:   renderUsecase,
	}
}
