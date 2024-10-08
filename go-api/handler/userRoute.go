package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swanhtetaungphyo/burmaguru/dto"
	"github.com/swanhtetaungphyo/burmaguru/services"
	"github.com/swanhtetaungphyo/burmaguru/utils"
	"strconv"
)

func GetAllUsers(c *fiber.Ctx) error {
	userJsonArr, err := services.GetAllUser()
	if err != nil {
		return utils.ErrorResponse(c, "User retrieval failed", fiber.ErrBadRequest.Code)
	}
	return utils.SuccessResponse(c, "Users retrieval is successful", 200, userJsonArr)
}

func GetUser(c *fiber.Ctx) error {
	requestParam := c.Params("id")
	userId, err := strconv.Atoi(requestParam)
	if err != nil {
		return utils.ErrorResponse(c, "Invalid parameter", fiber.ErrBadRequest.Code)
	}

	mediumUser, err := services.GetUserByIdService(int64(userId))
	if err != nil {
		return utils.ErrorResponse(c, err.Error(), fiber.ErrInternalServerError.Code)
	}

	return utils.SuccessResponse(c, "User retrieval by ID is successful", 200, mediumUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(c, "Invalid user ID", fiber.ErrBadRequest.Code)
	}
	var updateUserDto dto.UserDto
	if err := c.BodyParser(&updateUserDto); err != nil {
		return utils.ErrorResponse(c, "Invalid request body", fiber.ErrBadRequest.Code)
	}
	updatedUser, err := services.UpdateUserService(int64(id), updateUserDto)
	if err != nil {
		return utils.ErrorResponse(c, "Failed to update user", fiber.ErrInternalServerError.Code)
	}

	return utils.SuccessResponse(c, "User updated successfully", 200, updatedUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(c, "Invalid user ID", fiber.ErrBadRequest.Code)
	}
	user, err := services.DeleteUserService(int64(id))
	if err != nil {
		return utils.ErrorResponse(c, "Failed to delete user", fiber.ErrInternalServerError.Code)
	}
	return utils.SuccessResponse(c, "User deleted successfully", 200, user)
}
