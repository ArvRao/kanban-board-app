package handler

import (
	"github.com/arvrao/kanban-board-app/database"
	"github.com/arvrao/kanban-board-app/model"
	"github.com/gofiber/fiber/v2"
)

// get all projects
func GetAllProjects(c *fiber.Ctx) error {
	db := database.DB.Db
	var projects []model.Project
	// find all projects in the database
	db.Find(&projects)
	if len(projects) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No projects available right now", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Projects for given user is found", "data": projects})
}
