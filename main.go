package main

import (
	"github.com/SarathLUN/golang-fiber-sqlite/controllers"
	"github.com/SarathLUN/golang-fiber-sqlite/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"net/http"
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
		AllowOrigins:     "*",
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
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Welcome to Golang, Fiber, SQLite, and GORM",
			"status":  "success",
		})
	})

	// start server
	err := app.Listen("localhost:8000")
	if err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}
