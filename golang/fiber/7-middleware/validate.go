package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(ValidateTokenMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome, valid user!")
	})

	log.Fatal(app.Listen(":8080"))
}

func ValidateTokenMiddleware(c *fiber.Ctx) error {
	token := c.Get("X-Auth-Token")

	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "X-Auth-Token header is missing",
		})
	} else if token != "my-secret-token" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	return c.Next()
}

/*

curl localhost:8080 -X GET
{"error":"X-Auth-Token header is missing"}

curl localhost:8080 -X GET -H 'X-Auth-Token: salam'
{"error":"Invalid token"}

curl localhost:8080 -X GET -H 'X-Auth-Token: my-secret-token'
Welcome, valid user!

*/
