package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/services"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)

func GetAllUsers(c *fiber.Ctx) error { 
	userJsonArr, err := services.GetAllUser()	
	if err != nil {
		return utils.ErrorResponse(c, "User retrieval failed", fiber.ErrBadRequest.Code)
	}
	return utils.SuccessResponse(c, "Users retrieval is successful", 200, userJsonArr)
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