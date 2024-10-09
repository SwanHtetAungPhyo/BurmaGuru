package handler

import "github.com/gofiber/fiber/v2"

func SetUp(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/countries", GetCountryHandler)
	api.Post("/countries/:id", CreateCountryHandler)
	api.Put("/countries/:id", UpdateCountryHandler)
	app.Delete("/countries", DeleteCountryHandler)
}
