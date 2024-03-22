package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/rarepepi/g1/config"
	"github.com/rarepepi/g1/database"
	_ "github.com/rarepepi/g1/docs"
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

// @title G1 GoLang Fiber API Template
// @version 1.0
// @description This is a simple API template using GoLang and Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
    // Connect to the Database
    database.ConnectDB()

    // Start a new fiber app
    app := fiber.New()

    // Swagger Docs
	app.Get("/swagger/*", swagger.HandlerDefault) // default
    app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))


    // Setup routes
    setUpRoutes(app)

    // Setup CORS
    app.Use(cors.New())

    // 404 Handler
    app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Set port to os env or default to 3000
	port := "3000"
	if osPort := config.Config("PORT"); osPort != "" {
		port = fmt.Sprintf("%s", osPort)
	}

    // Listen on PORT 3000
	log.Fatal(app.Listen("0.0.0.0:" + port))
}