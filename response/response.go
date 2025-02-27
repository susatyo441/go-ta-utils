package response

import "github.com/gofiber/fiber/v2"

// SendResponse is a helper to send JSON responses in Fiber
func SendResponse(c *fiber.Ctx, statusCode int, data interface{}, message string) error {
	response := fiber.Map{
		"status":  statusCode, //changed attribute from "statusCode" to "status" (following TS nestjs app)
		"message": message,
		"data":    data,
	}
	return c.Status(statusCode).JSON(response)
}

// SUCCESS
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return SendResponse(c, fiber.StatusOK, data, message)
}

func Created(c *fiber.Ctx, message string, data interface{}) error {
	return SendResponse(c, fiber.StatusCreated, data, message)
}

// ERROR
func BadRequest(c *fiber.Ctx, message string, data interface{}) error {
	return SendResponse(c, fiber.StatusBadRequest, data, message)
}

func Unauthorized(c *fiber.Ctx, message string, data interface{}) error {
	return SendResponse(c, fiber.StatusUnauthorized, data, message)
}

func Forbidden(c *fiber.Ctx, message string, data interface{}) error {
	return SendResponse(c, fiber.StatusForbidden, data, message)
}

func NotFound(c *fiber.Ctx, message string, data interface{}) error {
	return SendResponse(c, fiber.StatusNotFound, data, message)
}

func InternalServerError(c *fiber.Ctx, message string, data interface{}) error {
	return SendResponse(c, fiber.StatusInternalServerError, data, message)
}
