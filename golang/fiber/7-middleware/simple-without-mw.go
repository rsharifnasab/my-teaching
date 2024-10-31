package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		start := time.Now()

		time.Sleep(1 * time.Second)
		err := c.SendString("hello fiber")

		end := time.Since(start)
		println(end.Milliseconds())

		return err
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.SendString("about page")

		end := time.Since(start)
		println(end.Milliseconds())

		return err
	})

	log.Fatal(app.Listen(":8080"))
}
