package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/search", func(c *fiber.Ctx) error {
		query := c.Query("q", "default")
		return c.SendString("Search query: " + query)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		println(string(c.Body()))
		return c.SendString("Hello, I got your message!")
	})

	log.Fatal(app.Listen(":8080"))
}

// installation
// register handler on path+method
// function signature: context and error
// how "error" is used?
// curl 'http://localhost:8080/' -i
