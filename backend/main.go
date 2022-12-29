package main

import (
	"fmt"

	"github.com/arvrao/kanban-board-app/router"
	"github.com/arvrao/kanban-board-app/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)
	fmt.Println("Running on port 3000")
	app.Listen(":3000")
}
