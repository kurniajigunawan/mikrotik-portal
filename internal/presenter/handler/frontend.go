package handler

import (
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Render(c fiber.Ctx) error {
	page := c.Params("page")
	err := c.Render(page, nil)
	if err != nil {
		return c.Render("maint/not_found", nil)
	}
	return nil
}
