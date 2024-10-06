package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/models"
	"github.com/swanhtetaungphyo/burmaguru/services"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)

func Register(c *fiber.Ctx) error {
	var inputUser models.User
	if err := c.BodyParser(&inputUser); err != nil{
		return utils.ErrorResponse(c, "Invalide input", fiber.StatusBadRequest)
	}
	registeredUser, err := services.RegistrationServices(&inputUser); 
	if err != nil{
		return utils.ErrorResponse(c, "failed to registered",fiber.ErrBadRequest.Code)
	}
	return utils.SuccessResponse(c, "User registered successfully",200, registeredUser)
}

func VerifyEmail(c * fiber.Ctx) error {
	return  utils.SuccessResponse(c,"Verification successful", 200, nil)

}


func Login(c *fiber.Ctx) error {
	return utils.SuccessResponse(c,"Login successful", 200, nil)
}