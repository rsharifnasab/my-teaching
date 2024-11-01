package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{}))

	// Or extend your config for customization
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			// return c.IP() == "127.0.0.1"
			return false
		},

		Max: 2,

		Expiration: 30 * time.Second,

		KeyGenerator: func(c *fiber.Ctx) string {
			// default: c.IP
			// return c.Get("x-forwarded-for")
			return c.IP()
		},

		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		println(c.IP())
		return c.SendString("ok")
	})

	log.Fatal(app.Listen(":8080"))
}
