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

func ShowPlans(c *fiber.Ctx) error {
	return c.Render("pages/plans", fiber.Map{
		"Title": "plans",
	})
}

func ShowContact(c *fiber.Ctx) error {
	return c.Render("pages/contact", fiber.Map{
		"Title": "Contact Us",
	})
}

func ShowFeatures(c *fiber.Ctx) error {
	return c.Render("pages/features", fiber.Map{
		"Title": "Features",
	})
}
