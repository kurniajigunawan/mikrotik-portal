package handler

import (
	mikrotikUC "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/mikrotik"
)

type Handler struct {
	mikrotikUsecase mikrotikUC.UsecaseItf
}

// NewHandler is a function to create a new handler
func NewHandler(mikrotikUsecase mikrotikUC.UsecaseItf) *Handler {
	return &Handler{
		mikrotikUsecase: mikrotikUsecase,
	}
}
