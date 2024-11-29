package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/watevadafa/retrospec/internal/routes"
	"log"
)

func main() {
	// Initialize template engine
	engine := html.New("./templates", ".html")

	// Initialize fiber app
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "main",
	})

	// Serve static files
	app.Static("/assets", "./assets")

	// Setup routes
	routes.SetupWebRoutes(app)
	routes.SetupAPIRoutes(app)

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Error handling
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).Render("errors/404", fiber.Map{
			"Title": "Page Not Found",
		})
	})

	// Start server
	log.Fatal(app.Listen(":8080"))
}
