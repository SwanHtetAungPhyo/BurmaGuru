package utils

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, message string, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"body":    data,
	})
}
func ErrorResponse(c *fiber.Ctx, message string, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "error",
		"message": message,
		"Error":   data,
	})
}
