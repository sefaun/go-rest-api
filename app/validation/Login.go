package validation

import (
	"go-rest-api/app/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginType struct {
	Username string `json:"username" validate:"required,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

func Login(c *fiber.Ctx) error {

	// Create a new login auth struct.
	login_type := &LoginType{}

	// Checking received data from JSON body.
	if err := c.BodyParser(login_type); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Create a new validator for a Login model.
	validate := utils.NewValidator()

	// Validate login_type fields.
	if err := validate.Struct(login_type); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": utils.ValidatorErrors(err),
		})
	}

	return c.Next()
}
