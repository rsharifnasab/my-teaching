package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recovermw "github.com/gofiber/fiber/v2/middleware/recover"
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

	app := fiber.New(fiber.Config{
		ErrorHandler: server.ErrorHandler,
	})

	v1 := app.Group("/v1")
	v1.Use(logger.New())
	v1.Use(recovermw.New())

	server.Register(v1)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":8080"))
}
