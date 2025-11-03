package render

import "github.com/gofiber/fiber/v3"

type GetPageResponse struct {
	Page string
	Data fiber.Map
}
