package app

import (
	"cringe-link-shortener/internal/postgres"
	"log"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	web    *fiber.App
	errLog *log.Logger
	db     postgres.DBHandler
}

func CreateApp(db postgres.DBHandler, errLog *log.Logger) *App {
	application := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.SendString(err.Error())
		},
	})

	result := &App{
		web:    application,
		errLog: errLog,
		db:     db,
	}

	setRoutes(result)

	return result
}

func (app *App) Run(addr string) {
	app.errLog.Fatalln(app.web.Listen(addr))
}

func (app *App) Shutdown() error {
	return app.web.Shutdown()
}
