package mainroute

import (
	apphandller "github.com/ekbal41/fiberX/main/handler"
	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router) {
	app.Get("/", apphandller.HomePage)
	app.Get("/about", apphandller.AboutPage)
	app.Get("/blog", apphandller.BlogPage)
	app.Get("/contact", apphandller.ContactPage)
	app.Get("/services", apphandller.ServPage)
}