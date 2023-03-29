package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"github.cpm/ekbal41/fiber-test-1/route"
)
func setupRoutes(app *fiber.App) {

	route.HomeRoutes(app.Group("/"))
	apiHome := app.Group("/api")
	route.TodoRoutes(apiHome.Group("/todo"))
	apiHome.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("You are in the api endpoint!")

	})

	

}
//go:embed views/*
var viewsfs embed.FS
//go:embed static/*
var staticfs embed.FS

func main(){
	// viewEngine := html.New("./views", ".html")
	viewEngine := html.NewFileSystem(http.FS(viewsfs), ".html")
	viewEngine.Reload(true)
	viewEngine.Debug(true)
	app := fiber.New(
		fiber.Config{
			Views: viewEngine,

		},
	)
    app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.FS(staticfs),
		PathPrefix: "static",
		Browse: true,
}))
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
