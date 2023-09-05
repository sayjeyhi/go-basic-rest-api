package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sayjeyhi.com/todolist/models"
)

// DbConnection is the database connection
var (
	DbConnection *gorm.DB
)

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=postgres port=5432"

	DbConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected! 🎉")
	log.Println("Running migrations 🚀")

	errConnection := DbConnection.AutoMigrate(&models.Todo{})
	if errConnection != nil {
		return
	}

	log.Println("Migration did run successfully 🎉")
}
