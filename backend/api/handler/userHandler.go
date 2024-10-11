package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaGuru/dto"
	"github.com/swanhtetaungphyo/burmaGuru/services"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"strconv"
)

func GetUser(ctx *fiber.Ctx) error {
	var userFromService dto.UserDTO
	userId, _ := strconv.Atoi(ctx.Params("id"))
	var err error
	if userFromService, err = services.GetUser(userId); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Faild to retrieval", err.Error())
	}
	return utils.ApiResponse(ctx, 200, "Retrieval  success", userFromService)
}

// CreateUser TODO: Delete after Authentication set up
func CreateUser(ctx *fiber.Ctx) error {
	var requestedUser dto.UserDTO
	if err := ctx.BodyParser(&requestedUser); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Invalid Request Format", err.Error())
	}
	if err := services.CreateUser(requestedUser); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "Error in User creation", err.Error())
	}
	return utils.ApiResponse(ctx, 200, "User created", "")
}

//TODO: write after the auth set up
//func UpdateUser(ctx *fiber.Ctx) error {
//
//
//}

func DeleteUser(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Params(":id"))
	if err := services.DeleteUser(userId); err != nil {
		return utils.ApiResponse(ctx, fiber.StatusInternalServerError, "User Removal failed", err)
	}
	return utils.ApiResponse(ctx, 200, "User Removal succeed", "")
}
