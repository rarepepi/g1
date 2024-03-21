package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rarepepi/g1/database"
	"github.com/rarepepi/g1/router"
)

func main() {
    // Start a new fiber app
    app := fiber.New()

    // Find .env file
    err := godotenv.Load(".env")
    if err != nil{
    log.Fatalf("Error loading .env file: %s", err)
    }

    // Connect to the Database
    database.ConnectDB()

    // Setup the router
    router.SetupRoutes(app)

    // Listen on PORT 3000
    app.Listen(":3000")
}