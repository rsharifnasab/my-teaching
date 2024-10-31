package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{}))

	app.Get("/v1", func(c *fiber.Ctx) error {
		return c.SendString("Hello, I got your message!")
	})

	log.Fatal(app.Listen(":8080"))
}
