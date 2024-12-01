package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/watevadafa/retrospec/internal/api/auth"
)

func SetupAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")

	// Authentication routes
	authGroup := apiGroup.Group("/auth")
	authGroup.Post("/signup", auth.SignUp)
	// auth.Post("/login", api.Login)

	// // Protected routes
	// apiGroup.Use(api.AuthMiddleware)
}
