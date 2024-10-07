package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)


func CreateGuide(c *fiber.Ctx) error {
	
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}

func GetGuide(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}

func UpdateGuide(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}

func DeleteGuide(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}