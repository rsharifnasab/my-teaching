package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// mutable shared state
var global int

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		global = global + 1

		result := fmt.Sprintf("number: %d\n", global)
		return c.SendString(result)
	})

	log.Fatal(app.Listen(":8080"))
}
