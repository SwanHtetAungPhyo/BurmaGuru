package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func SuccessResponse(c *fiber.Ctx, message string, statusCode int, data interface{}) error {
	response := &Response{
		Status:  statusCode,
		Message: message,
		Body:    data,
	}
	return c.Status(statusCode).JSON(response)
}

func ErrorResponse(c *fiber.Ctx, message string, statusCode int, data interface{}) error {
	response := &Response{
		Status:  statusCode,
		Message: message,
		Body:    data,
	}
	return c.Status(statusCode).JSON(response)
}
