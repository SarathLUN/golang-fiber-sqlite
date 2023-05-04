package controllers

import "github.com/gofiber/fiber/v2"

func LoadHomePage(c *fiber.Ctx) error {
	// Render a template named 'home.page.html' with content
	return c.Render("home.page", fiber.Map{
		"Title":       "Home",
		"Description": "This is home page",
	})
}
