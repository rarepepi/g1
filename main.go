package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rarepepi/g1/database"
	"github.com/rarepepi/g1/routes"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", routes.Hello)
	app.Get("/allbooks", routes.AllBooks)
	app.Post("/addbook", routes.AddBook)
	app.Post("/book", routes.Book)
	app.Put("/update", routes.Update)
	app.Delete("/delete", routes.Delete)
}


func main() {
    // Connect to the Database
    database.ConnectDB()

    // Start a new fiber app
    app := fiber.New()

    // Setup routes
    setUpRoutes(app)

    // Setup CORS
    app.Use(cors.New())

    // 404 Handler
    app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

    // Listen on PORT 3000
	log.Fatal(app.Listen(":3000"))
}