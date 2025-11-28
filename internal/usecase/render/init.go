package render

import (
	"context"

	"github.com/gofiber/fiber/v3"
	homePageRepo "github.com/kurniajigunawan/mikrotik-portal/internal/repository/homepage"
	serviceRepo "github.com/kurniajigunawan/mikrotik-portal/internal/repository/service"
)

type RenderMapItf interface {
	Render(ctx context.Context) fiber.Map
}

type UsecaseItf interface {
	GetPage(ctx context.Context, page string) (GetPageResponse, error)
}

type Usecase struct {
	homepageRepo homePageRepo.Interface
	serviceRepo  serviceRepo.Interface
}

func New(homepageRepo homePageRepo.Interface, serviceRepo serviceRepo.Interface) UsecaseItf {
	return &Usecase{
		homepageRepo: homepageRepo,
		serviceRepo:  serviceRepo,
	}
}
