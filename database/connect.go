package database

import (
	"log"

	"github.com/rarepepi/g1/config"
	"github.com/rarepepi/g1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// ConnectDB connect to db
func ConnectDB() {
	log.Println("🐵 Connecting to Database 🐵")	
    
    // Connection URL to connect to Postgres Database
	dsn := config.Config("DATABASE_URL")
 	
	// Connect to the DB and initialize the DB variable
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("🐵 Connected to Database! 🐵")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("🐵 Running DB migrations... 🐵")
	db.AutoMigrate(&models.Book{})

	DB = Dbinstance{
		Db: db,
	}
}