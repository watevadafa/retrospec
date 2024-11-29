package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	// Initialize template engine
	engine := html.New("./frontend/templates", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "main",
	})

	// Serve static files
	app.Static("/assets", "./assets")

	// Basic route to test the template
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/home", fiber.Map{
			"Title": "Home",
		})
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
