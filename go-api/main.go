package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
	"github.com/swanhtetaungphyo/burmaguru/database"
	"github.com/swanhtetaungphyo/burmaguru/handler"
	"github.com/swanhtetaungphyo/burmaguru/utils"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.ConnectDB()

	logFile := utils.LogInit()
	if logFile == nil {
		fmt.Println("Log initialization failed. Exiting program.")
		return
	}
	defer func(logfile *os.File) {
		err := logfile.Close()
		if err != nil {
			return
		}
	}(logFile)

	app.Use(limiter.New(limiter.Config{
		Max:        5,
		Expiration: 60 * 1000 * 1000 * 1000,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.IP()
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			return utils.ErrorResponse(ctx, "You have reached the API call limit", fiber.StatusTooManyRequests, "")
		},
	}))
	routes.Setup(app)

	log.Println("Server is starting...")
	fmt.Println("API is running on http://localhost:8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
