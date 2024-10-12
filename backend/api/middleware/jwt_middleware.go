package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/swanhtetaungphyo/burmaGuru/config"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
)

func JwtProtected() fiber.Handler {
	return jwtware.New(
		jwtware.Config{
			SigningKey:   []byte(config.GetJWTSecret()),
			ContextKey:   "user",
			ErrorHandler: JwtError,
		})
}

func JwtError(ctx *fiber.Ctx, err error) error {
	if err != nil {
		return utils.ApiResponse(ctx, fiber.StatusUnauthorized, "Unauthorized", err.Error())
	}
	return nil
}
