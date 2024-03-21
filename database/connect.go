package database

import (
	"fmt"
	"log"

	"github.com/rarepepi/g1/config"
	"github.com/rarepepi/g1/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
    var err error
	fmt.Println("Connecting to Database")	
    if err != nil {
        log.Println("Idiot")
    }

    // Connection URL to connect to Postgres Database
	dsn := config.Config("DATABASE_URL")
 	
	// Connect to the DB and initialize the DB variable
	DB, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	DB.AutoMigrate(&model.Note{})
	fmt.Println("Database Migrated")
}