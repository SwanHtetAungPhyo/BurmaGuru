package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaGuru/dto"
	"github.com/swanhtetaungphyo/burmaGuru/services"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
)

func RecommendationHandler(ctx *fiber.Ctx) error {
	var preference dto.UserPreference
	if err := ctx.BodyParser(&preference); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusBadRequest, "Invalid Data format", err.Error())
	}

	recommendations, err := services.Recommender(preference)
	if err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Error", err)
	}

	return utils.ApiResponse(ctx, fiber.StatusCreated, "Recommender System succssed", dto.RecommendationArray{Recommendation: recommendations})

}
