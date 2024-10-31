package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Middleware to log request method and path
func logMethodAndPath(c *fiber.Ctx) error {
	log.Printf("1\n")
	err := c.Next()
	log.Printf("5\n")
	return err
}

func logHeaders(c *fiber.Ctx) error {
	log.Printf("2\n")
	err := c.Next()
	log.Printf("4\n")
	return err
}

func main() {
	app := fiber.New()
	app.Use(logHeaders)                     // print 2 - 4
	app.Use(logMethodAndPath)               // print 1 - 5
	app.Get("/", func(c *fiber.Ctx) error { // print 3
		log.Printf("3\n")
		return c.SendString("Hello, Fiber with Middleware Chaining!")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		println("i am handler")
		return c.SendString("About page with Middleware Chaining")
	})

	log.Fatal(app.Listen(":8080"))
}
