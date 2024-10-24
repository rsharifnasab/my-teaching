package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8080"))
}

// installation
// register handler on path+method
// function signature: context and error
// how "error" is used?
// curl 'http://localhost:8080/' -i
