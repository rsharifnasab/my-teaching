package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("mysecretkey")

func main() {
	app := fiber.New()

	app.Post("/login", login)

	// Protected routes group
	app.Use("/api", jwtware.New(jwtware.Config{
		SigningKey: jwtSecret,
	}))

	// A protected route that requires a valid JWT token
	app.Get("/api/dashboard", dashboard)

	log.Fatal(app.Listen(":3000"))
}

func login(c *fiber.Ctx) error {
	// In a real app, you would validate user credentials (e.g., username, password)

	// Create JWT claims
	claims := jwt.MapClaims{
		"username": "john_doe",                            // Add any custom claims you want
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

func dashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to the protected dashboard!",
	})
}

// curl -X POST http://localhost:3000/login | jq '.token' -r
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA1MzYxNzYsInVzZXJuYW1lIjoiam9obl9kb2UifQ.Zu2TAQmXJGgViRmyR3y59ND2HTTB9FtXvL-b4ASCluc

// curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA1MzYxNzYsInVzZXJuYW1lIjoiam9obl9kb2UifQ.Zu2TAQmXJGgViRmyR3y59ND2HTTB9FtXvL-b4ASCluc" http://localhost:3000/api/dashboard
// {"message":"Welcome to the protected dashboard!"}
