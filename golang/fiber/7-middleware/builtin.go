package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	// example of nested groups
	groupLogs := app.Group("/logs")
	group := groupLogs.Group("/inner")

	group.Use(logger.New(logger.Config{}))
	group.Use(recover.New())

	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	app.Get("/panic",
		// recover.New(), // register MW for only one handler
		func(c *fiber.Ctx) error {
			panic("something went wrong!")
		},
	)

	log.Fatal(app.Listen(":8080"))
}
