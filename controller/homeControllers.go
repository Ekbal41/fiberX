package controller

import "github.com/gofiber/fiber/v2"

func HomePage(c *fiber.Ctx) error {
	return c.Render("views/index", fiber.Map{
		"Title": "Home Page",
	},"views/layouts/main")
}
func AboutPage(c *fiber.Ctx) error {
	return c.Render("views/about", fiber.Map{
		"Title": "About Page",
	},"views/layouts/main")
}