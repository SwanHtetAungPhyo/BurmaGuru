package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaGuru/dto"
	"github.com/swanhtetaungphyo/burmaGuru/services"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
)

func RegisterHandler(ctx *fiber.Ctx) error {
	var inputUser dto.UserDTO
	if err := ctx.BodyParser(&inputUser); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Invalid data format", err)
	}
	if err := services.Registration(&inputUser); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to register", err)
	}
	return utils.ApiResponse(ctx, 200, "User registration is successful", "")
}

func LoginHandler(ctx *fiber.Ctx) error {
	loginData := new(dto.LoginRequest) // Use dto.LoginRequest

	if err := ctx.BodyParser(loginData); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Cannot parse body", nil)
	}

	accessToken, refreshToken, err := services.Login(loginData)
	if err != nil {
		return utils.ApiResponse(ctx, fiber.StatusUnauthorized, "Login failed", nil)
	}

	return utils.ApiResponse(ctx, fiber.StatusOK, "Login successful", fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func VerifyEmail(ctx *fiber.Ctx) error {
	token := ctx.Query("token")
	if token == "" {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Token for email verification is left", "")
	}
	return utils.ApiResponse(ctx, 200, "Email verification is successful", "")
}
