package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"sayjeyhi.com/todolist/models"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("Health check ✅")
	})

	api := app.Group("/api", logger.New())
	v1 := api.Group("/v1")
	v1.Get("/todos", models.GetTodos)
	v1.Get("/todos/:id", models.GetTodo)
	v1.Post("/todos", models.CreateTodo)
	v1.Delete("/todos/:id", models.DeleteTodo)
	v1.Patch("/todos/:id", models.UpdateTodo)
}
