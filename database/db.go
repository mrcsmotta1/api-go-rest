package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB // Using GORM for ORM functionality
	err error
)

func Connect() {
	// Initialize the database connection here
	// This is a placeholder; actual implementation will depend on your database setup
	// For example, you might use gorm.Open() to connect to a specific database
	// DB, err = gorm.Open("your_database_driver", "your_connection_string")
	stringConnection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(stringConnection))
	if err != nil {
		log.Panic("Failed to connect to the database")
	}
}
