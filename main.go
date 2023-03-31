package main

import (
	"embed"
	"log"
	"net/http"

	apiroute "github.com/ekbal41/fiberX/v1/api/route"
	approute "github.com/ekbal41/fiberX/v1/app/route"
	"github.com/ekbal41/fiberX/v1/detabase"
	"github.com/ekbal41/fiberX/v1/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"github.com/spf13/viper"
)

//runs when initializing the app --------------------------->
func init() {
	detabase.ConnectDB()
	detabase.InitMigration()
}

//setupRoutes sets up all the routes for the application---->
func setupRoutes(app *fiber.App) {
	approute.Routes(app.Group("/"))
	apiHome := app.Group("/api")
	apiroute.Routes(apiHome.Group("/"))
}




//returns custom 404 page--------------------------->
func notFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).Render("views/404", fiber.Map{})

}



//all static file setup---------------------------->
//go:embed views/*
var viewsfs embed.FS
//go:embed static/*
var staticfs embed.FS



//main function------------------------------------>
func main(){
	//load config file
	utils.LoadConfig()
	//setup view engine
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
	log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}
