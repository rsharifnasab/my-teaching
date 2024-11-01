package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Simulated action that requires authentication
	app.Post("/transfer2", func(c *fiber.Ctx) error {
		// Check if the user is authenticated (i.e., has a session cookie)
		if c.Cookies("session_id") != "123456" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		fmt.Printf("c: %+v\n", c)

		// Simulate transferring money
		amount := c.FormValue("amount")
		return c.SendString(fmt.Sprintf("Transferred $%s!", amount))
	})

	// Or extend your config for customization
	app.Use(csrf.New(csrf.Config{
		KeyLookup: "form:csrf_token", // Token is expected in request headers

		CookieName:     "csrf_token",
		CookieSameSite: "Strict",

		Expiration:   1 * time.Hour,
		KeyGenerator: utils.UUIDv4,
		ContextKey:   "csrf_token", // Custom key in Locals for accessing the CSRF token

	}))

	app.Get("/csrf", func(c *fiber.Ctx) error {
		csrfToken := fmt.Sprintf("%s", c.Locals("csrf_token"))
		return c.SendString("CSRF Token: " + csrfToken)
	})

	// Simulate a "login" route
	app.Get("/login", func(c *fiber.Ctx) error {
		// Set a session cookie to simulate login
		c.Cookie(&fiber.Cookie{
			Name:     "session_id",
			Value:    "123456", // Simulate a session ID
			Path:     "/",
			SameSite: "None", // Allow this cookie to be sent with cross-origin requests
		})
		return c.SendString("Logged in successfully!")
	})

	// Simulated action that requires authentication
	app.Post("/transfer", func(c *fiber.Ctx) error {
		// Check if the user is authenticated (i.e., has a session cookie)
		if c.Cookies("session_id") != "123456" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		fmt.Printf("c: %+v\n", c)

		// Simulate transferring money
		amount := c.FormValue("amount")
		return c.SendString(fmt.Sprintf("Transferred $%s!", amount))
	})

	// Serve a simple HTML form for transferring money
	app.Get("/", func(c *fiber.Ctx) error {
		csrfToken := c.Locals("csrf_token")
		fmt.Printf("generated csrf: %v\n", csrfToken)

		if csrfToken == nil {
			return c.Status(fiber.StatusInternalServerError).SendString("CSRF token is missing!")
		}
		html := fmt.Sprintf(`
			<h1>Transfer Money</h1>
			<form action="/transfer" method="POST">
				<input type="hidden" name="csrf_token" value="%s">
				<label for="amount">Amount:</label>
				<input type="text" id="amount" name="amount" value="100">
				<button type="submit">Transfer</button>
			</form>
		`, csrfToken)
		c.Response().Header.Add("Content-Type", "text/html")
		return c.Send([]byte(html))
	})

	log.Fatal(app.Listen(":3000"))
}
