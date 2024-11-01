package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			return ctx.Status(code).SendString(err.Error())
		},
	})

	app.Use(logger.New())

	app.Get("/search", func(c *fiber.Ctx) error {
		return fmt.Errorf("i cannot search")
	})
	app.Get("/", func(c *fiber.Ctx) error {
		println(string(c.Body()))
		return c.SendString("Hello, I got your message!")
	})

	log.Fatal(app.Listen(":8080"))
}
