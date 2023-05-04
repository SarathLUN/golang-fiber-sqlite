package main

import (
	"github.com/SarathLUN/golang-fiber-sqlite/controllers"
	"github.com/SarathLUN/golang-fiber-sqlite/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html" // add engine
	"log"
)

func init() {
	// connect to database via initializes package
	initializers.ConnectDB()
}
func main() {
	// Initialize a standard Go's html/template engine
	engine := html.New("./views", ".html")

	// Create a new Fiber template with template engine
	web := fiber.New(fiber.Config{
		Views: engine,
	})
	// serve static files
	web.Static("/static", "./public")
	// define routes to load home page from controllers package
	web.Get("/", controllers.LoadHomePage)

	// start server
	err := web.Listen("localhost:3000")
	if err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}
