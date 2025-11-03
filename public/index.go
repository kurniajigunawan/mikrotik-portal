package public

import (
	"github.com/gofiber/fiber/v3"
	component "github.com/kurniajigunawan/mikrotik-portal/public/component"
	"github.com/kurniajigunawan/mikrotik-portal/public/style"
	widget "github.com/kurniajigunawan/mikrotik-portal/public/widget"
)

var PrimaryColor = style.ColorIndigo500

var HomePage = fiber.Map{
	"Heading": widget.Heading{
		Title:    "Hotspot Portal",
		Subtitle: "Troubleshoot your Hotspot Account device with ease.",
	},
	"MenuItem": []widget.MenuItem{
		{
			LinkURL:   "/reset",
			Title:     "Reset Sessions",
			Subtitle:  "Remove all active sessions based on account. All connected devices will be disconnected.",
			Icon:      "fas fa-redo-alt",
			IconColor: "green-400",
		},
	},
}

var ResetPage = fiber.Map{
	"Heading": widget.Heading{
		Title:    "Reset Sessions",
		Subtitle: "Remove your account active sessions. Please relogin all devices after reset sessions.",
	},
	"Form": widget.Form{
		ActionURL: "/api/reset-session",
		Method:    "POST",
		Input: []widget.Input{
			{
				ID:    "username",
				Label: "Username",
				Name:  "username",
				Type:  "text",
			},
		},
		SubmitButton: component.ButtonSolid{
			Text:                 "Submit",
			Type:                 "submit",
			BackgroundColor:      PrimaryColor,
			BackgroundHoverColor: PrimaryColor.ChangeStep(-1),
		},
	},
	"BackButton": component.ButtonLink{
		Text:           "Back to Home",
		LinkURL:        "/home",
		TextColor:      style.ColorIndigo400,
		TextHoverColor: style.ColorIndigo300,
		Icon:           "fa-solid fa-arrow-left",
	},
}
