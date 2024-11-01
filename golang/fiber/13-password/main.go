package main

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var userStore = struct {
	sync.RWMutex
	store map[string]string
}{
	store: make(map[string]string),
}

func main() {
	app := fiber.New()

	app.Post("/register", func(c *fiber.Ctx) error {
		type RegisterRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var req RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request format",
			})
		}

		// [salt][hash(password + pepper + salt)]
		// [salt][hash(password + salt)]
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to hash password",
			})
		}

		userStore.Lock()
		userStore.store[req.Username] = string(hashedPassword)
		userStore.Unlock()

		return c.JSON(fiber.Map{
			"message": "User registered successfully",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		type LoginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request format",
			})
		}

		userStore.RLock()
		hashedPassword, exists := userStore.store[req.Username]
		userStore.RUnlock()

		if !exists {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid username or password",
			})
		}

		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid username or password",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Login successful",
		})
	})

	// Start Fiber server
	log.Fatal(app.Listen(":3000"))
}
