package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/swanhtetaungphyo/burmaGuru/config"
	"github.com/swanhtetaungphyo/burmaGuru/db"
	"github.com/swanhtetaungphyo/burmaGuru/middleware"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"log"
	"net"
	"os"
)

func main() {
	applicaion := fiber.New()

	config.LoadEnv()

	if err := db.SetUpDataBase(); err != nil {
		log.Printf("Error in %v is %v", utils.CurrentFunction(), err.Error())
		return
	}

	db.Migrate()

	logFile := utils.LogInit()
	if logFile == nil {
		log.Printf("LogFile initializaion is Failed ......")
		return
	}
	defer func(logfile *os.File) {
		if err := logfile.Close(); err != nil {
			log.Printf("Error is in %v and reason is :%v ", utils.CurrentFunction(), err.Error())
			return
		}
	}(logFile)
	applicaion.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 60 * 1000 * 1000 * 1000,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.IP()
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			return utils.ApiResponse(ctx, fiber.StatusTooManyRequests, "API call reaches the limit", net.Interface{})
		},
	}))

	applicaion.Use("/auth", middleware.JwtProtected())

	config.SetUp(applicaion)

	applicaion.Listen(":8080").Error()
}
