package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
)

func Logger(c *fiber.Ctx) error {
	fmt.Println("Logger middleware executed")
	return c.Next()
}

func main() {
	app := fiber.New()

	// app.Use(Logger)
	app.Use(
		skip.New(
			Logger,
			func(ctx *fiber.Ctx) bool { // predicate
				return ctx.Method() == fiber.MethodGet
			},
		),
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("response get")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("response post")
	})

	app.Listen(":8080")
}

// curl localhost:8080/ -i -X GET
// curl localhost:8080/ -i -X POST
