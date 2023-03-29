package route

import (
	"github.com/gofiber/fiber/v2"
	"github.cpm/ekbal41/fiber-test-1/controller"
)

func TodoRoutes(route fiber.Router) {
	route.Get("/", controller.Welcome )
}