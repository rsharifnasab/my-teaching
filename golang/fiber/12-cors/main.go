package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	apiApp := fiber.New()
	apiApp.Use(logger.New())
	apiApp.Use(cors.New(cors.Config{
		// allow xdomain in y backend
		AllowOrigins: "http://localhost:3000,http://127.0.0.1:3000",
		AllowMethods: "GET,POST",
	}))
	apiApp.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("This is data from the API server. (GET)")
	})
	apiApp.Post("/api", func(c *fiber.Ctx) error {
		return c.SendString("This is data from the API server. (POST)")
	})
	go func() {
		// ydomain (destination): localhost:4000
		log.Fatal(apiApp.Listen(":4000"))
	}()

	clientApp := fiber.New()
	apiApp.Use(logger.New())
	clientApp.Static("/", "./client")
	log.Fatal(clientApp.Listen(":3000"))
	// xdomain (request source): 127.0.0.1:3000
}
