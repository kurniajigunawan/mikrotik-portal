package render

import "github.com/gofiber/fiber/v3"

const (
	HomePage = "home"
)

type GetPageResponse struct {
	Page string
	Data fiber.Map
}
