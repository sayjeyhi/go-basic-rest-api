package models

import (
	"github.com/gofiber/fiber/v2"
	"sayjeyhi.com/todolist/database"
)

type Todo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DbConnection
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DbConnection
	var todo Todo
	err := db.Find(&todo, id)
	if err != nil {
		return c.Status(500).SendString("Something went wrong")
	}
	return c.JSON(&todo)
}

func CreateTodo(c *fiber.Ctx) error {
	db := database.DbConnection
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	db.Create(&todo)
	return c.JSON(&todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DbConnection
	var todo Todo
	db.First(&todo, id)
	if todo.Title == "" {
		return c.Status(500).SendString("No todo found with ID")
	}
	db.Delete(&todo)
	return c.SendString("Todo successfully deleted")
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DbConnection
	var todo Todo
	db.First(&todo, id)
	if todo.Title == "" {
		return c.Status(500).SendString("No todo found with ID")
	}
	if err := c.BodyParser(&todo); err != nil {
		return err
	}
	db.Save(&todo)
	return c.JSON(&todo)
}
