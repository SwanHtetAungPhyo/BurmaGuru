package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)


func CreateComment(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}

func GetComments(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}


func DeleteComment(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}
