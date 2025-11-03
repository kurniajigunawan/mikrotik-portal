package render

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/kurniajigunawan/mikrotik-portal/public"
)

type UsecaseItf interface {
	GetPage(ctx context.Context, page string) (GetPageResponse, error)
}

type Usecase struct {
	pages map[string]fiber.Map
}

func New() UsecaseItf {
	return &Usecase{
		pages: map[string]fiber.Map{
			"reset": public.ResetPage,
			"home":  public.HomePage,
		},
	}
}
