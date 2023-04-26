package presentation

import (
	"github.com/gofiber/fiber/v2"

	"fiberpgx1/business"
)

type Presenter struct {
	Admin *business.Admin
}

func (p *Presenter) AdminCreateArticle(ctx *fiber.Ctx) error {
	in := business.CreateArticleIn{}
	err := ctx.BodyParser(&in)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	out := p.Admin.CreateArticle(in)
	if out.Err != nil {
		// TODO: differentiate by type of error
		return ctx.Status(500).JSON(fiber.Map{
			"error": out.Err.Error(),
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"in":  in,
		"out": out,
	})
}
