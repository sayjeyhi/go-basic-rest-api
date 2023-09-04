package database

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DbConnection is the database connection
var (
	DbConnection *gorm.DB
)
