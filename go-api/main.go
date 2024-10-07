package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/databases"
	"github.com/swanhtetaungphyo/burmaguru/handler"
	"github.com/swanhtetaungphyo/burmaguru/utils"
	"github.com/joho/godotenv"
)

func main() {
    app := fiber.New()

    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

	databases.ConnectDB()
   
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
