package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/services"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"strconv"
)

func CreateCountryHandler(ctx *fiber.Ctx) error {
	var newCountry models.Country
	if err := ctx.BodyParser(&newCountry); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Invalid data", err)
	}
	if err := services.CreateCountry(newCountry); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to create country", err)
	}
	return utils.ApiResponse(ctx, fiber.StatusCreated, "Country created successfully", newCountry)
}

func UpdateCountryHandler(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Invalid country ID", err)
	}

	var updatedCountry models.Country
	if err := ctx.BodyParser(&updatedCountry); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Invalid data", err)
	}

	if err := services.UpdateCountry(uint(id), updatedCountry); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to update country", err)
	}

	return utils.ApiResponse(ctx, fiber.StatusOK, "Country updated successfully", updatedCountry)
}

func GetCountryHandler(ctx *fiber.Ctx) error {
	countries, err := services.GetAllCountry()
	if err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve countries", err)
	}
	return utils.ApiResponse(ctx, fiber.StatusOK, "Countries retrieved successfully", countries)
}

func GetCountryByIdHandler(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	country, err := services.GetCountryById(uint(id))
	if err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve country", err)
	}
	return utils.ApiResponse(ctx, fiber.StatusOK, "Country retrieved successfully", country)
}

func DeleteCountryHandler(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	if err := services.DeleteCountry(uint(id)); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to delete country", err)
	}
	return utils.ApiResponse(ctx, fiber.StatusOK, "Country deleted successfully", nil)
}
