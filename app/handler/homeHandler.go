package apphandller

import (
	"github.com/gofiber/fiber/v2"
)


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
func BlogPage(c *fiber.Ctx) error {
	return c.Render("views/blog", fiber.Map{
		"Title": "Blog Page",
	},"views/layouts/main")
}
func ContactPage(c *fiber.Ctx) error {
	return c.Render("views/contact", fiber.Map{
		"Title": "Contact Page",
	},"views/layouts/main")
}
func ServPage(c *fiber.Ctx) error {
	return c.Render("views/services", fiber.Map{
		"Title": "Service Page",},"views/layouts/main")
}