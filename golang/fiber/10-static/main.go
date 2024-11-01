package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Serve static files from a directory (e.g. "public")
	// This will serve any files in the "public" directory at the root URL
	app.Static("/", "./public")

	// Optional: You can also specify a different route for serving static files
	// Example: Serve static files from "/assets" route
	app.Static("/assets", "./assets")

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
