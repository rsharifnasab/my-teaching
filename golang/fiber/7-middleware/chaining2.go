package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Middleware to log request method and path
func logMethodAndPath2(c *fiber.Ctx) error {
	log.Printf("> Method: %s, Path: %s\n", c.Method(), c.Path())

	err := c.Next()
	log.Printf("< Method: %s, Path: %s\n", c.Method(), c.Path())

	return err
}

func logHeaders2(c *fiber.Ctx) error {
	log.Printf("> Headers: %v\n", c.GetReqHeaders())

	err := c.Next()
	log.Printf("< Headers: %v\n", c.GetReqHeaders())

	return err
}

func main() {
	app := fiber.New()

	app.Use(logMethodAndPath2) // first and last
	app.Use(logHeaders2)

	app.Get("/", func(c *fiber.Ctx) error {
		log.Printf("handling\n")
		return c.SendString("Hello, Fiber with Middleware Chaining!")
	})

	log.Fatal(app.Listen(":8080"))
}
