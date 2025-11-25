package render

import (
	"context"

	homePageRepo "github.com/kurniajigunawan/mikrotik-portal/internal/repository/homepage"
	serviceRepo "github.com/kurniajigunawan/mikrotik-portal/internal/repository/service"
)

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
