package render

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/kurniajigunawan/mikrotik-portal/internal/usecase/render/template/formpage"
	"github.com/kurniajigunawan/mikrotik-portal/internal/usecase/render/template/homepage"
)

func (u *Usecase) GetPage(ctx context.Context, page string) (GetPageResponse, error) {
	var data fiber.Map
	switch page {
	case HomePage:
		var err error
		data, page, err = homepage.Build(u.homepageRepo).Render(ctx).Compile()
		if err != nil {
			return GetPageResponse{}, fmt.Errorf("failed to render homepage: %v", err)
		}
	case "reset":
		var err error
		data, page, err = formpage.Build(u.serviceRepo).Render(ctx).Compile()
		if err != nil {
			return GetPageResponse{}, fmt.Errorf("failed to render form: %v", err)
		}
	default:
		return GetPageResponse{}, fmt.Errorf("page %s not found", page)
	}
	return GetPageResponse{
		Page: page,
		Data: data,
	}, nil
}
