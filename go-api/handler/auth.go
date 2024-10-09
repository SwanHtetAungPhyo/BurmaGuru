package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/dto"
	"github.com/swanhtetaungphyo/burmaguru/services"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)

//type LoginParameter struct {
//	Email    string `json:"Email"`
//	Password string `json:"Password"`
//}
//

func Register(c *fiber.Ctx) error {
	var inputUser dto.UserDto
	if err := c.BodyParser(&inputUser); err != nil {
		return utils.ErrorResponse(c, "Invalide input", fiber.StatusBadRequest, err.Error())
	}
	if err := services.RegistrationServices(&inputUser); err != nil {
		return utils.ErrorResponse(c, "failed to registered", fiber.ErrBadRequest.Code, err.Error())
	}
	return utils.SuccessResponse(c, "User registered successfully", 200, "")
}

//
//func VerifyEmail(c *fiber.Ctx) error {
//	token := c.Query("token")
//	if token == "" {
//		return utils.ErrorResponse(c, "Email verification failed", fiber.ErrBadRequest.Code, "")
//	}
//
//	return utils.SuccessResponse(c, "Verification successful", 200, nil)
//
//}
//
//func Login(c *fiber.Ctx) error {
//	var inputLogin LoginParameter
//
//	if err := c.BodyParser(&inputLogin); err != nil {
//		return utils.ErrorResponse(c, "Invalid Input", fiber.ErrBadRequest.Code, err.Error())
//	}
//
//	email, password := inputLogin.Email, inputLogin.Password
//
//	log.Printf("Login attempt: Email=%s, Password=****", email)
//	if err := services.LoginService(email, password); err != nil {
//		return utils.ErrorResponse(c, "Login Failed: "+err.Error(), fiber.ErrUnauthorized.Code, err.Error())
//	}
//
//	return utils.SuccessResponse(c, "Login successful", fiber.StatusOK, nil)
//}
