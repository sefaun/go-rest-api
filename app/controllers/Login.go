package controllers

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {

	user := "sefa createds"

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": false,
		"data":    user,
	})
}
