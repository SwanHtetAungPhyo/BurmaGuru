package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/routes"
)

func main(){
	app := fiber.New()

	routes.Setup(app)
	fmt.Printf("API is running on the port, http://localhost:8080")
	app.Listen(":8080")
}