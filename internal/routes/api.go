package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/watevadafa/retrospec/internal/api"
)

func SetupAPIRoutes(app *fiber.App) {
	// apiGroup := app.Group("/api")

	// Authentication routes
	// auth := apiGroup.Group("/auth")
	// auth.Post("/register", api.Register)
	// auth.Post("/login", api.Login)

	// // Protected routes
	// apiGroup.Use(api.AuthMiddleware)
}
