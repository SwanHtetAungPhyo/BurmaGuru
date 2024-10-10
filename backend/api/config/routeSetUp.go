package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaGuru/handler"
)

func SetUp(app *fiber.App) {
	api := app.Group("/api")

	countries := api.Group("/countries")
	countries.Get("/", handler.GetCountryHandler)
	countries.Post("/", handler.CreateCountryHandler)
	countries.Put("/id", handler.UpdateCountryHandler)
	countries.Delete("/:id", handler.DeleteCountryHandler)

	recommendation := api.Group("/recomendation")
	recommendation.Post("/", handler.RecommendationHandler)

	universities := countries.Group("/:country_id/universities")
	universities.Get("/", handler.GetUniversities)
	universities.Post("/", handler.CreateUniversityHandler)
	universities.Put("/:id", handler.UpdateUniversity)

	resources := countries.Group("/:country_id/resources")
	resources.Post("/", handler.CreateResources)
	resources.Get("/", handler.GetResources)
	resources.Put("/:resource_id", handler.UpdateResource)
	resources.Delete("/:resource_id", handler.DeleteResource)

}
