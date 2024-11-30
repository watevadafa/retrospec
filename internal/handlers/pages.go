package handlers

import "github.com/gofiber/fiber/v2"

func ShowHome(c *fiber.Ctx) error {
	return c.Render("pages/home", fiber.Map{
		"Title": "Home",
	})
}

func ShowAbout(c *fiber.Ctx) error {
	return c.Render("pages/about", fiber.Map{
		"Title": "About",
	})
}

func ShowPricing(c *fiber.Ctx) error {
	return c.Render("pages/pricing", fiber.Map{
		"Title": "Pricing",
	})
}

func ShowContact(c *fiber.Ctx) error {
	return c.Render("pages/contact", fiber.Map{
		"Title": "Contact Us",
	})
}
