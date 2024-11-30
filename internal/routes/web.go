package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/watevadafa/retrospec/internal/handlers"
)

func SetupWebRoutes(app *fiber.App) {
	app.Get("/", handlers.ShowHome)
	app.Get("/home", handlers.ShowHome)
	app.Get("/about", handlers.ShowAbout)
	app.Get("/pricing", handlers.ShowPricing)
	app.Get("/contact", handlers.ShowContact)
	app.Get("/signup", handlers.ShowSignup)
	app.Get("/signup/pro", handlers.ShowSignupPro)
}
