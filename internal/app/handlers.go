package app

import (
	"cringe-link-shortener/pkg/translator"

	"github.com/gofiber/fiber/v2"
)

// goTo - возвращает исконную ссылку по сокращению.
//
// Принимает: контекст.
//
// Возвращает: ошибку.

// @Summary      Returns link by the shortened.
// @Description  Get link added to db by the key
// @Tags         links
// @Accept       plain
// @Produce      plain
// @Param        key path string true "Key for the link"
// @Success      200 {string} string "Link"
// @Failure      400 {string} string "Error message"
// @Failure      504 {string} string "Error message"
// @Router       / [get]
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

// addLink - добавляет ссылку в БД.
//
// Принимает: контекст.
//
// Возвращает: ошибку.

// @Summary      Adds input link to the DB.
// @Description  Add link to db and get its key.
// @Tags         links
// @Accept       json
// @Produce      plain
// @Param        link body string true "Input link"
// @Success      200 {string} string "Key"
// @Failure      400 {string} string "Error message"
// @Failure      500 {string} string "Error message"
// @Router       / [post]
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
