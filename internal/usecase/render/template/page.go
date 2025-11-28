package page

import (
	"context"

	"github.com/gofiber/fiber/v3"
)

type Render struct {
	err  error
	data fiber.Map
	slug string
}

func CreateRender(slug string, data fiber.Map, err error) Render {
	return Render{
		slug: slug,
		err:  err,
		data: data,
	}
}

func (r Render) Compile() (fiber.Map, string, error) {
	return r.data, r.slug, r.err
}

type Interface interface {
	Render(ctx context.Context) Render
}
