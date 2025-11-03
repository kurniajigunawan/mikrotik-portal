package handler

import (
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Render(c fiber.Ctx) error {
	page := c.Params("page")
	response, err := h.renderUsecase.GetPage(c.Context(), page)
	if err != nil {
		return c.Render("maint/not_found", nil)
	}
	return c.Render(response.Page, response.Data)
}
