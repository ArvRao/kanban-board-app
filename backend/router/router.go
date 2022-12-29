package router

import "github.com/gofiber/fiber/v2"

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/todo")
	// routes
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("Hello there @arvind")
	})
}
