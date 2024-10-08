package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/models"
	"github.com/swanhtetaungphyo/burmaguru/services"
	"github.com/swanhtetaungphyo/burmaguru/utils"
	"strconv"
)

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return utils.ErrorResponse(c, "Failed to parse request body", 400, err.Error())
	}

	createdCategory, err := services.CreateCategoryService(category)
	if err != nil {
		return utils.ErrorResponse(c, "Failed to create category", 500, err.Error())
	}

	return utils.SuccessResponse(c, "Category created successfully", 201, createdCategory)
}

func GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrorResponse(c, "Failed to retrieve category", 401, err.Error())
	}
	category, err := services.GetCategoryByIdService(categoryId)
	if err != nil {
		return utils.ErrorResponse(c, "Failed to retrieve category", 500, err.Error())
	}

	return utils.SuccessResponse(c, "Category retrieved successfully", 200, category)
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return utils.ErrorResponse(c, "Failed to parse request body", 400, err.Error())
	}

	var err error
	category.ID, err = strconv.Atoi(id)
	if err != nil {
		return utils.ErrorResponse(c, "Failed to Cast the type", 400, err.Error())
	}

	if err := services.UpdateCategoryService(category); err != nil {
		return utils.ErrorResponse(c, "Failed to update category", 500, err.Error())
	}

	return utils.SuccessResponse(c, "Category updated successfully", 200, category)
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	intId, err := StringToInt(id)
	if err != nil {
		return utils.ErrorResponse(c, "Failed to make conversion", 500, err.Error())
	}
	if err := services.DeleteCategoryService(intId); err != nil {
		return utils.ErrorResponse(c, "Failed to delete category", 500, err.Error())
	}

	return utils.SuccessResponse(c, "Category deleted successfully", 200, nil)
}

func GetCategories(c *fiber.Ctx) error {
	categories, err := services.GetAllCategoriesService()
	if err != nil {
		return utils.ErrorResponse(c, "Failed to retrieve categories", 500, err.Error())
	}

	return utils.SuccessResponse(c, "Categories retrieved successfully", 200, categories)
}

func StringToInt(input string) (int, error) {
	return strconv.Atoi(input)
}
