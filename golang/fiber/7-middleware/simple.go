package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func customMiddleware(c *fiber.Ctx) error {
	log.Printf("Request: %s %s\n", c.Method(), c.Path())

	start := time.Now()
	defer func() {
		duration := time.Since(start)
		log.Printf("Processed in %v", duration)
	}()

	err := c.Next()

	return err
}

func main() {
	app := fiber.New()

	app.Use(customMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(1 * time.Second)
		return c.SendString("Hello, Fiber!")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("About page")
	})

	log.Fatal(app.Listen(":8080"))
}
