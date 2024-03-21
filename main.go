package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rarepepi/g1/database"
	"github.com/rarepepi/g1/router"
)

func main() {
    // Start a new fiber app
    app := fiber.New()

    // Connect to the Database
    database.ConnectDB()

    // Setup the router
    router.SetupRoutes(app)

    // Listen on PORT 3000
    app.Listen(":3000")
}