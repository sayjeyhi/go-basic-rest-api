package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sayjeyhi.com/todolist/database"
	"sayjeyhi.com/todolist/models"
)

func initDB() {
	var err error
	dsn := "host=localhost user=postgres password=postgres port=5432"

	database.DbConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected! 🎉")
	log.Println("Running migrations 🚀")

	errConnection := database.DbConnection.AutoMigrate(&models.Todo{})
	if errConnection != nil {
		return
	}

	log.Println("Migration did run successfully 🎉")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("Health check ✅")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/todos", models.GetTodos)
	v1.Get("/todos/:id", models.GetTodo)
	v1.Post("/todos", models.CreateTodo)
	v1.Delete("/todos/:id", models.DeleteTodo)
	v1.Patch("/todos/:id", models.UpdateTodo)
}

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(cors.New())

	initDB()
	setUpRoutes(app)

	log.Fatal(app.Listen(":3010"))
}
