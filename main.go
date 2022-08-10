package main

import (
	"go-rest-api/configs"
	"go-rest-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	config := configs.FiberConfig()

	app := fiber.New(config)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api := app.Group("/api/v1")

	routes.AllRoutes(api)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
		})
	})

	app.Listen(":5000")
}
