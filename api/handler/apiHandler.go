package apihandler

import (
	"github.com/gofiber/fiber/v2"
)


func Home(c *fiber.Ctx) error {
	return c.SendString("This is api endpoint")
}