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
// @Accept		 plain
// @Produce      json
// @Param        key path string true "Key for the link"
// @Success      302 {string} string "Redirects to link"
// @Header       302 {string} Location "Url to redirect"
// @Failure      400 {object} outErr "Error message"
// @Failure      404 {object} outErr "Error message"
// @Router       /{key} [get]
func (app *App) goTo(c *fiber.Ctx) error {
	shortened := c.Params("key")
	id, err := translator.Translate(shortened)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(getErr(err))
	}
	url, err := app.db.GetLink(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(getErr(err))
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
// @Produce      json
// @Param        link body input true "Input link"
// @Success      200 {object} output "Key"
// @Failure      400 {object} outErr "Error message"
// @Failure      500 {object} outErr "Error message"
// @Router       / [post]
func (app *App) addLink(c *fiber.Ctx) error {
	body := input{}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(getErr(err))
	}
	link, err := getUrl(body.Link)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(getErr(err))
	}
	id, err := app.db.AddLink(link)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(getErr(err))
	}
	shortened := translator.Encrypt(id)

	return c.JSON(output{shortened})
}
