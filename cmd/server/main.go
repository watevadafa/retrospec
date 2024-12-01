package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	"github.com/watevadafa/retrospec/internal/config"
	"github.com/watevadafa/retrospec/internal/db"
	"github.com/watevadafa/retrospec/internal/handlers"
	"github.com/watevadafa/retrospec/internal/models/repositories"
	"github.com/watevadafa/retrospec/internal/routes"
	"log"
)

func main() {
	// Initialize database
	if err := db.InitDB(config.GetDBPath()); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.DB.Close() // Don't forget to close the DB connection

	// Single line to initialize all repositories
	repositories.InitializeRepositories(db.DB)

	// Initialize template engine
	engine := html.New("./templates", ".html")

	// Initialize fiber app
	app := fiber.New(fiber.Config{
		Views:        engine,
		ViewsLayout:  "main",
		ErrorHandler: handlers.ErrorHandler,
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

	// Start server
	log.Fatal(app.Listen(":8080"))
}
