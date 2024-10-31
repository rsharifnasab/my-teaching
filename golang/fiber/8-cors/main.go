package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	apiApp := fiber.New()
	apiApp.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "POST",
	}))
	apiApp.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("This is data from the API server.")
	})

	go func() {
		log.Fatal(apiApp.Listen(":4000"))
	}()

	clientApp := fiber.New()

	clientApp.Static("/", "./client")

	log.Fatal(clientApp.Listen(":3000"))
}
