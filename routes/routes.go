package routes

import (
	"go-rest-api/app/controllers"
	"go-rest-api/app/validation"

	"github.com/gofiber/fiber/v2"
)

func AllRoutes(api fiber.Router) {

	api.Post("/login", validation.Login, controllers.Login)

}
