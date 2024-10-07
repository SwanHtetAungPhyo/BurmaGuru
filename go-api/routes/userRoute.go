package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)


func GetAllUsers(c *fiber.Ctx) error { 
	
	return utils.SuccessResponse(c, "Users reterival is successful", 200, nil)
}

func GetUser(c *fiber.Ctx) error {

	return utils.SuccessResponse(c, "User reterival by ID is successful", 200, nil)

}

func UpdateUser(c *fiber.Ctx) error{
	return utils.SuccessResponse(c, "User is updated succcessfully", 200, nil)
}

func DeleteUser(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"User is deleted successfulsly", 200, nil)
}