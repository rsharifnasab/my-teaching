package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func CustomTimeoutMiddleware(duration time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), duration)
		defer cancel()

		done := make(chan error, 1)

		go func() {
			done <- c.Next() // Call the next middleware/handler
		}()

		select {
		case err := <-done:
			return err
		case <-ctx.Done():
			log.Println("Request timed out in custom middleware")
			return c.Status(fiber.StatusRequestTimeout).SendString("Custom Timeout Error: Request took too long!")
		}
	}
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// deprecated
	app.Get("/slow1",
		timeout.New(
			func(c *fiber.Ctx) error {
				time.Sleep(3 * time.Second)
				return c.SendString("This should timeout!")
			},
			2*time.Second,
		),
	)

	app.Use(CustomTimeoutMiddleware(2 * time.Second))
	app.Get("/slow2", func(c *fiber.Ctx) error {
		time.Sleep(3 * time.Second)
		return c.SendString("This is slow 2")
	})

	log.Fatal(app.Listen(":8080"))
}
