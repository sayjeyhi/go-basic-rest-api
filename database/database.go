package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sayjeyhi.com/todolist/config"
	"sayjeyhi.com/todolist/models"
	"strconv"
)

// DbConnection is the database connection
var (
	DbConnection *gorm.DB
)

func InitDB() {
	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s sslmode=disable TimeZone=%s",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_TIMEZONE"),
	)

	DbConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", config.Config("DB_NAME"))
	DbConnection.Exec(createDatabaseCommand)

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
