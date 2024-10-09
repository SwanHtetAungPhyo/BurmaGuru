package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/swanhtetaungphyo/burmaGuru/db"
	"github.com/swanhtetaungphyo/burmaGuru/handler"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"log"
	"net"
)

func main() {
	applicaion := fiber.New()

	if err := db.SetUpDataBase(); err != nil {
		log.Printf("Error in %v is %v", utils.CurrentFunction(), err.Error())
		return
	}

	db.Migrate()

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

	handler.SetUp(applicaion)

	applicaion.Listen(":8080").Error()
}
