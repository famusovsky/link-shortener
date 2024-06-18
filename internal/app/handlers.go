package app

import (
	"cringe-link-shortener/pkg/translator"

	"github.com/gofiber/fiber/v2"
)

func (app *App) goTo(c *fiber.Ctx) error {
	shortened := c.Params("key")
	id, err := translator.Translate(shortened)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	url, err := app.db.GetLink(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.Redirect(url)
}

func (app *App) addLink(c *fiber.Ctx) error {
	body := struct {
		Link string
	}{}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	id, err := app.db.AddLink(body.Link)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	shortened := translator.Encrypt(id)

	return c.SendString(shortened)
}
