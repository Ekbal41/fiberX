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

// -------------------setupRoutes sets up all the routes for the application------------------>
func setupRoutes(app *fiber.App) {
	route.HomeRoutes(app.Group("/"))
	apiHome := app.Group("/api")
	route.TodoRoutes(apiHome.Group("/todo"))
	apiHome.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("You are in the api endpoint!")

	})
}






// ----------------------------------- returns custom 404 page--------------------------------->
func notFound(c *fiber.Ctx) error {
    // return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	// or render view
	return c.Status(fiber.StatusNotFound).Render("views/404", fiber.Map{})

}





//---------------------------------------all static file setup---------------------------->
//go:embed views/*
var viewsfs embed.FS
//go:embed static/*
var staticfs embed.FS




//-------------------------------main function----------------------------------------->
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
	// app.Static("/static", "./static")
    app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.FS(staticfs),
		PathPrefix: "static",
		Browse: true,
}))
	setupRoutes(app)
	//this should be after all the routes
	app.Use(notFound)
	log.Fatal(app.Listen(":3000"))
}
