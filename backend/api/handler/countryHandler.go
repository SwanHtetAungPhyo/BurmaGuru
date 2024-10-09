package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/services"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
)

func CreateCountryHandler(ctx *fiber.Ctx) error {
	return utils.ApiResponse(ctx, fiber.StatusCreated, "Country is created successfully", utils.EmptyMap())
}

func UpdateCountryHandler(ctx *fiber.Ctx) error {
	return utils.ApiResponse(ctx, fiber.StatusCreated, "Country is created successfully", utils.EmptyMap())
}

func GetCountryHandler(ctx *fiber.Ctx) error {
	var countries []models.Country
	countries, err := services.GetAllCountry()
	if err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Country retrival failed", err)
	}
	return utils.ApiResponse(ctx, fiber.StatusCreated, "Country is created successfully", countries)
}

func GetCountryByIdHandler(ctx *fiber.Ctx) error {
	return utils.ApiResponse(ctx, fiber.StatusCreated, "Country is created successfully", utils.EmptyMap())
}
func DeleteCountryHandler(ctx *fiber.Ctx) error {
	return utils.ApiResponse(ctx, fiber.StatusCreated, "Country is created successfully", utils.EmptyMap())
}
