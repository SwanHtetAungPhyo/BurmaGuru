package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)

func CreateArticle(c *fiber.Ctx) error {
	
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}

func GetArticle(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}

func UpdateArticle(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}

func DeleteArticle(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}


func GetArticles(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Successful Function", 200, nil)
}