package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)


func CreateCategory(c *fiber.Ctx) error {
	
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}

func GetCategory(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}

func UpdateCategory(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}

func DeleteCategory(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}


func GetCategories(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}