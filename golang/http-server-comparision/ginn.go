package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Basic route
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Gin Framework!")
	})

	// Route with query parameters (e.g., /greet?name=John)
	router.GET("/greet", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + name,
		})
	})

	// Route with path parameters (e.g., /user/123)
	router.GET("/user/:id", func(c *gin.Context) {
		userID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"user_id": userID,
		})
	})

	// POST route with JSON body (e.g., {"username": "john", "password": "pass123"})
	router.POST("/login", func(c *gin.Context) {
		var json struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.Username == "admin" && json.Password == "admin" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	})

	// Grouped routes (e.g., /api/v1/*)
	api := router.Group("/api/v1")
	{
		api.GET("/time", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"time": time.Now().Format(time.RFC3339),
			})
		})
	}

	// Start the server
	router.Run(":8080")
}

/*
curl -X GET http://localhost:8080/
curl -X GET "http://localhost:8080/greet?name=John"
curl -X GET http://localhost:8080/user/123
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"username": "admin", "password": "admin"}'
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"username": "user", "password": "wrong_pass"}'

curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"username": "admin"}'

curl -X GET http://localhost:8080/api/v1/time

*/
