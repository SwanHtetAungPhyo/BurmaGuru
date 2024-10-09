package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/services"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"strconv"
)

func CreateUniversityHandler(ctx *fiber.Ctx) error {
	var university models.University

	if err := ctx.BodyParser(&university); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Invalid data", err)
	}

	if err := services.CreateUniversity(university); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to create university", err)
	}

	return utils.ApiResponse(ctx, fiber.StatusCreated, "University created successfully", university)
}
func UpdateUniversity(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var updatedUniversity models.University
	if err := ctx.BodyParser(&updatedUniversity); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Invalid data", err)
	}
	if err := services.UniUpdateService(uint(id), updatedUniversity); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to update university", err)
	}
	return utils.ApiResponse(ctx, fiber.StatusOK, "University updated successfully", nil)
}

func GetUniversities(ctx *fiber.Ctx) error {

	countryID, err := strconv.Atoi(ctx.Params("country_id"))
	if err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Invalid country ID", err)
	}

	universities, err := services.GetUniversities(uint(countryID))
	if err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve universities", err)
	}

	return utils.ApiResponse(ctx, fiber.StatusOK, "Universities retrieved successfully", universities)
}
