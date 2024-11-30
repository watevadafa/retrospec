package handlers

import "github.com/gofiber/fiber/v2"

func ShowHome(c *fiber.Ctx) error {
	return c.Render("pages/home", fiber.Map{
		"Title": "Home",
	})
}
