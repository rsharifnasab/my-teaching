package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler for the root route
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handler to demonstrate query parameters
func getUser(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	return c.String(http.StatusOK, "Name: "+name+", Age: "+age)
}

// Handler for path parameters
func getUserByID(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"ID": id,
	})
}

// Handler for JSON request and response
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

func main() {
	e := echo.New()

	// Routes
	e.GET("/", hello)               // Basic route
	e.GET("/user", getUser)         // Query parameter route
	e.GET("/user/:id", getUserByID) // Path parameter route
	e.POST("/user", createUser)     // JSON post route

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

/*
curl http://localhost:8080/
curl "http://localhost:8080/user?name=John&age=30"
curl http://localhost:8080/user/123
curl -X POST http://localhost:8080/user -d '{"name":"Alice","age":25}' -H "Content-Type: application/json"
*/
