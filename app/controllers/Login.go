package controllers

import (
	"go-rest-api/app/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	fields := make(map[string]interface{})

	fields["username"] = "sefa"
	fields["password"] = "mypassword"
	fields["token"] = "BearerToken"

	return utils.SuccessResponse(c, "user_created", fields)
}
