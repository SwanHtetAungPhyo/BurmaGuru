package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/databases"
	"github.com/swanhtetaungphyo/burmaguru/routes"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)

func main() {
	app := fiber.New()

	err := databases.InitDB("main.db")
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	logFile := utils.LogInit()
	if logFile == nil {
		fmt.Println("Log initialization failed. Exiting program.")
		return
	}
	defer logFile.Close()

	routes.Setup(app)

	log.Println("Server is starting...")

	fmt.Println("API is running on http://localhost:8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}


