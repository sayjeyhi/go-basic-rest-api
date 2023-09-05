package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"sayjeyhi.com/todolist/handlers"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("Health check ✅")
	})

	api := app.Group("/api", logger.New())
	v1 := api.Group("/v1")

	// Auth
	auth := v1.Group("/auth")
	auth.Post("/login", handlers.Login)

	// Todos
	v1.Get("/todos", handlers.GetTodos)
	v1.Get("/todos/:id", handlers.GetTodo)
	v1.Post("/todos", handlers.CreateTodo)
	v1.Delete("/todos/:id", handlers.DeleteTodo)
	v1.Patch("/todos/:id", handlers.UpdateTodo)
}
