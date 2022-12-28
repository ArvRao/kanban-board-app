package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello there")
}

func main() {
	app := fiber.New()
	app.Get("/hello", Hello)
	fmt.Println("Running on port 3000")
	app.Listen(":3000")
}
