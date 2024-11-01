package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/session/v2"
)

var store = session.New()

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/login", loginHandler)
	app.Get("/dashboard", authRequired, dashboardHandler)
	app.Get("/logout", logoutHandler)

	log.Fatal(app.Listen(":3000"))
}

func loginHandler(c *fiber.Ctx) error {
	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var creds Credentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	if creds.Username == "admin" && creds.Password == "password" {
		sess := store.Get(c)

		sess.Set("username", creds.Username)
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save session"})
		}

		return c.JSON(fiber.Map{"message": "login successful"})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
}

func dashboardHandler(c *fiber.Ctx) error {
	sess := store.Get(c)

	username := sess.Get("username")
	if username == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Welcome to your dashboard, %s!", username),
	})
}

func logoutHandler(c *fiber.Ctx) error {
	// Get the session
	sess := store.Get(c)

	if err := sess.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to destroy session"})
	}

	return c.JSON(fiber.Map{"message": "logout successful"})
}

func authRequired(c *fiber.Ctx) error {
	sess := store.Get(c)

	if sess.Get("username") == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	return c.Next()
}
