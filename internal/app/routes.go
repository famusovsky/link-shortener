package app

import "github.com/gofiber/swagger"

func setRoutes(app *App) {
	app.web.Get("/swagger/*", swagger.New())
	app.web.Get("/:key?", app.goTo)
	app.web.Post("/", app.addLink)
}
