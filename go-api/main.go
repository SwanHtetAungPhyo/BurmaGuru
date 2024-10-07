package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/swanhtetaungphyo/burmaguru/routes"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)

func main(){
	app := fiber.New()

	logFile := utils.LogInit()
	if logFile == nil {
		fmt.Println("Log initialization failed. Exiting program.")
		return
	}
	defer logFile.Close()

	routes.Setup(app)
	log.Info("Hello")
	fmt.Printf("API is running on the port, http://localhost:8080")
	app.Listen(":8080")
}


