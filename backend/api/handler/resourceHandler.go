package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/services"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"strconv"
)

func CreateResources(ctx *fiber.Ctx) error {
	var resource models.Resource
	if err := ctx.BodyParser(&resource); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Invalid request", err)
	}

	if err := services.CreateResource(resource); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, fmt.Sprintf("%v", err.Error()), err)
	}
	return utils.ApiResponse(ctx, fiber.StatusCreated, "Created Successfully", nil)
}

func GetResources(ctx *fiber.Ctx) error {
	countryId, _ := strconv.Atoi(ctx.Params("country_id"))
	resources, err := services.GetAllResources(uint(countryId))
	if err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Reterival Failed", err)
	}
	return utils.ApiResponse(ctx, fiber.StatusCreated, "Reterival Successed", resources)
}

func UpdateResource(ctx *fiber.Ctx) error {
	var resource models.Resource

	countryId, _ := strconv.Atoi(ctx.Params("country_id"))
	if err := ctx.BodyParser(&resource); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Invalid Data", err)
	}
	if err := services.UpdateResourcesSer(uint(countryId), resource); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Update Failed", err)
	}
	return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "UpdateScuccessed", "")
}

func DeleteResource(ctx *fiber.Ctx) error {
	countryId, err := strconv.Atoi(ctx.Params("country_id"))
	if err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Invalid country ID", err)
	}

	if err := services.DeleteResourceSer(uint(countryId)); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Failed to delete resources", err)
	}

	return utils.ApiResponse(ctx, fiber.StatusOK, "Resources deleted successfully", nil)
}
