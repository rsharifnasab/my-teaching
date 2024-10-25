package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm-postgres/database"
	"gorm-postgres/models"
	"gorm-postgres/repository"
	"gorm-postgres/routes"
)

func main() {
	db := database.ConnectDb()
	bookRepo := repository.NewGormBook(db)

	server := &routes.HTTPServer{
		Repo: bookRepo,
	}

	models.Migrate(db)

	app := fiber.New()
	app.Get(`/:test<int>`, func(c *fiber.Ctx) error {
		return c.SendString(c.Params("test"))
	})

	v1 := app.Group("/v1")

	server.Register(v1)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":8080"))
}
