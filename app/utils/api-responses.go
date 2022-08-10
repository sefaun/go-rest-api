package utils

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	fields := make(map[string]interface{})

	fields["success"] = true

	if message != "" {
		fields["message"] = message
	}

	if data != nil {
		fields["data"] = data
	}

	return c.Status(fiber.StatusOK).JSON(fields)
}

func ErrorResponse(c *fiber.Ctx, message string, data interface{}) error {
	fields := make(map[string]interface{})

	fields["success"] = false

	if message != "" {
		fields["message"] = message
	}

	if data != nil {
		fields["data"] = data
	}

	return c.Status(fiber.StatusBadRequest).JSON(fields)
}
