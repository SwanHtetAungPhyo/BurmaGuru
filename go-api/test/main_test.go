package test

import (
	"github.com/gofiber/fiber/v2"
	routes "github.com/swanhtetaungphyo/burmaguru/handler"
)

func SetApp() *fiber.App {
	app := fiber.New()

	app.Post("/auth/register", routes.Register)
	return app
}
