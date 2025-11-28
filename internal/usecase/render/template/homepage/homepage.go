package homepage

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/kurniajigunawan/mikrotik-portal/internal/repository/homepage"
	templateUsecase "github.com/kurniajigunawan/mikrotik-portal/internal/usecase/render/template"
	"github.com/kurniajigunawan/mikrotik-portal/public/widget"
)

type Usecase struct {
	slug string
	repo homepage.Interface
}

func Build(repo homepage.Interface) templateUsecase.Interface {
	return &Usecase{
		slug: "home",
		repo: repo,
	}
}

func (u *Usecase) Render(ctx context.Context) templateUsecase.Render {
	menus, err := u.repo.GetActiveMenu(ctx)
	if err != nil {
		return templateUsecase.CreateRender(u.slug, fiber.Map{}, err)
	}
	var menuItems []widget.MenuItem
	for _, menu := range menus {
		menuItems = append(menuItems, widget.MenuItem{
			LinkURL:   menu.Link,
			Title:     menu.Title,
			Subtitle:  menu.Description,
			Icon:      menu.Icon,
			IconColor: menu.IconColor,
		})
	}

	return templateUsecase.CreateRender(u.slug, fiber.Map{
		"Heading": widget.Heading{
			Title:    "Hotspot Portal",
			Subtitle: "Troubleshoot your Hotspot Account device with ease.",
		},
		"MenuItem": menuItems,
	}, nil)
}
