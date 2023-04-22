package main

import (
	"github.com/SarathLUN/golang-fiber-sqlite/controllers"
	"github.com/SarathLUN/golang-fiber-sqlite/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html" // add engine
	"log"
)

func init() {
	// connect to database via initializes package
	initializers.ConnectDB()
}
func main() {
	// create new fiber app
	app := fiber.New()
	// create another fiber app for api
	api := fiber.New()
	// mount api app to /api
	app.Mount("/api", api)
	// use logger middleware in the app
	app.Use(logger.New())
	// use cors to allow all origins
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	// register our routes handlers with the api app
	api.Route("/notes", func(r fiber.Router) {
		r.Get("", controllers.FindNotes)
		r.Post("/", controllers.CreateNote)
	})
	// create routes for /notes/:noteId
	api.Route("/notes/:noteId", func(r fiber.Router) {
		r.Get("", controllers.FindNoteById)
		r.Patch("", controllers.UpdateNote)
		r.Delete("", controllers.DeleteNote)
	})
	// define routes
	api.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Welcome to Golang, Fiber, SQLite, and GORM",
			"status":  "success",
		})
	})

	// Initialize a standard Go's html/template engine
	engine := html.New("./views", ".gohtml")

	// Create a new Fiber template with template engine
	web := fiber.New(fiber.Config{
		Views: engine,
	})
	// serve static files
	web.Static("/static", "./public")
	// define routes to load home page
	web.Get("/", func(c *fiber.Ctx) error {
		// Render a template named 'home.page.html' with content
		return c.Render("home.page", fiber.Map{
			"Title":       "Home",
			"Description": "This is home page",
		})
	})

	// start server
	err := web.Listen("localhost:3000")
	if err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}
