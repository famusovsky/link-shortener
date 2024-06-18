package app

func setRoutes(app *App) {
	app.web.Get("/:key?", app.goTo)
	app.web.Post("/", app.addLink)
}
