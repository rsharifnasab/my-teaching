package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	auth := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "password123", // Username: admin, Password: password123
		},
	})

	app.Use("/dashboard", auth)

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the protected dashboard!")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the public page!")
	})

	log.Fatal(app.Listen(":8080"))
}
