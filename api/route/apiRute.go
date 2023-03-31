package apiroute

import (
	apihandler "github.com/ekbal41/fiberX/v1/api/handler"
	"github.com/gofiber/fiber/v2"
)

func Routes(api fiber.Router) {
	api.Get("/", apihandler.Home)

	
}