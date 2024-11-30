package handlers

import "github.com/gofiber/fiber/v2"

func ShowSignup(c *fiber.Ctx) error {
	return c.Render("pages/signup", fiber.Map{
		"Title": "Sign Up",
		"IsPro": false,
	})
}

func ShowSignupPro(c *fiber.Ctx) error {
	return c.Render("pages/signup", fiber.Map{
		"Title": "Sign Up for Pro",
		"IsPro": true,
	})
}
