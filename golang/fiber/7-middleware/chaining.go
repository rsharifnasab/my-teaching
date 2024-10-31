package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Middleware to log request method and path
func logMethodAndPath(c *fiber.Ctx) error {
	log.Printf("Method: %s, Path: %s\n", c.Method(), c.Path())

	return c.Next()
}

func logHeaders(c *fiber.Ctx) error {
	log.Printf("Headers: %v\n", c.GetReqHeaders())

	return c.Next()
}

func main() {
	app := fiber.New()

	app.Use(logMethodAndPath)
	app.Use(logHeaders)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber with Middleware Chaining!")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("About page with Middleware Chaining")
	})

	log.Fatal(app.Listen(":8080"))
}
