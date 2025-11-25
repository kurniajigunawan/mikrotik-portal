package homepage

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/kurniajigunawan/mikrotik-portal/internal/repository/homepage"
	"github.com/kurniajigunawan/mikrotik-portal/public/widget"
)

type Interface interface {
	Render(ctx context.Context) fiber.Map
}

type Usecase struct {
	repo homepage.Interface
}

func Build(repo homepage.Interface) Interface {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Render(ctx context.Context) fiber.Map {
	return fiber.Map{
		"Heading": widget.Heading{
			Title:    "Hotspot Portal",
			Subtitle: "Troubleshoot your Hotspot Account device with ease.",
		},
		"MenuItem": u.renderMenuItem(ctx),
	}
}

func (u *Usecase) renderMenuItem(ctx context.Context) []widget.MenuItem {
	menus, err := u.repo.GetActiveMenu(ctx)
	if err != nil {
		return []widget.MenuItem{}
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
	return menuItems
}
