package routes

import (
	"github.com/SarathLUN/golang-fiber-sqlite/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"log"
)

func StartFrontEndRouter(baseURL string) {

	// Initialize a standard Go's html/template engine
	engine := html.New("./views", ".html")

	// Create a new Fiber template with template engine
	app := fiber.New(fiber.Config{
		Views:             engine,
		EnablePrintRoutes: true,
	})
	// use logger middleware
	app.Use(logger.New())
	// serve static files
	app.Static("/static", "./public")
	// define routes to load home page from controllers package
	app.Get("/", controllers.LoadHomePage)

	// start server
	err := app.Listen(baseURL)
	if err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}
