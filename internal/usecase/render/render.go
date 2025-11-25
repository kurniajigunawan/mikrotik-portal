package render

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/kurniajigunawan/mikrotik-portal/internal/usecase/render/homepage"
	"github.com/kurniajigunawan/mikrotik-portal/public/component"
	"github.com/kurniajigunawan/mikrotik-portal/public/style"
	"github.com/kurniajigunawan/mikrotik-portal/public/widget"
)

func (u *Usecase) GetPage(ctx context.Context, page string) (GetPageResponse, error) {
	var data fiber.Map
	if page == HomePage {
		data = homepage.Build(u.homepageRepo).Render(ctx)
	} else if page == "reset" {
		fields := []widget.FieldType{}
		service, err := u.serviceRepo.GetActiveServices(ctx)
		if err != nil {
			return GetPageResponse{}, fmt.Errorf("failed to get active services: %v", err)
		}
		options := make(map[string]string)
		for _, s := range service {
			options[fmt.Sprintf("%d", s.ID)] = s.Name
		}
		fields = append(fields, widget.Select{
			ID:      "service",
			Label:   "Service",
			Name:    "service",
			Options: options,
		}, widget.Input{
			ID:    "username",
			Label: "Username",
			Name:  "username",
			Type:  "text",
		})
		data = fiber.Map{
			"Heading": widget.Heading{
				Title:    "Reset Sessions",
				Subtitle: "Remove your account active sessions. Please relogin all devices after reset sessions.",
			},
			"Form": widget.Form{
				ActionURL: "/api/reset-session",
				Method:    "POST",
				Fields:    fields,
				SubmitButton: component.ButtonSolid{
					Text:            "Submit",
					Type:            "submit",
					BackgroundColor: style.ColorIndigo,
				},
			},
			"BackButton": component.ButtonLink{
				Text:      "Back to Home",
				LinkURL:   "/home",
				TextColor: style.ColorIndigo,
				Icon:      "fa-solid fa-arrow-left",
			},
		}
	} else {
		return GetPageResponse{}, fmt.Errorf("page %s not found", page)
	}
	return GetPageResponse{
		Page: page,
		Data: data,
	}, nil
}
