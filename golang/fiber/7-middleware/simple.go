package main

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func customMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	if start.Day() == 1 {
		return errors.New("i cannot work on 1st of each month")
	}

	err := c.Next()

	if false {
		// overrides the "next" response
		c.SendString("response in middleware")
	}

	duration := time.Since(start)
	log.Printf("Processed in %v", duration)

	return err
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(1 * time.Second)
		return c.SendString("Hello, Fiber!")
	})

	app.Use(customMiddleware)

	app.Get("/about", func(c *fiber.Ctx) error {
		println("about")
		return c.SendString("About page 2")
	})

	log.Fatal(app.Listen(":8080"))
}
