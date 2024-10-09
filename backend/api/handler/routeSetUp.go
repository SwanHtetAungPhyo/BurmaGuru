package handler

import "github.com/gofiber/fiber/v2"

func SetUp(app *fiber.App) {
	api := app.Group("/api")

	countries := api.Group("/countries")
	countries.Get("/", GetCountryHandler)
	countries.Post("/:id", CreateCountryHandler)
	countries.Put("/id", UpdateCountryHandler)
	countries.Delete("/:id", DeleteCountryHandler)

	universities := countries.Group("/:country_id/universities")
	universities.Get("/", GetUniversities)
	universities.Post("/", CreateUniversityHandler)
	universities.Put("/:id", UpdateUniversity)
}
