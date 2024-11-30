package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/watevadafa/retrospec/internal/handlers"
)

func SetupWebRoutes(app *fiber.App) {
	app.Get("/", handlers.ShowHome)
	// app.Get("/boards", handlers.ShowBoards)
	// app.Get("/boards/:id", handlers.ShowBoard)
	// app.Get("/organizations", handlers.ShowOrganizations)
}
