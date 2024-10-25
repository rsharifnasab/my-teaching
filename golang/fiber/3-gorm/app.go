package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm-postgres/database"
	"gorm-postgres/models"
	"gorm-postgres/routes"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/allbooks", routes.AllBooks)

	app.Post("/addbook", routes.AddBook)
	app.Post("/book", routes.Book)

	app.Put("/update", routes.Update)

	app.Delete("/delete", routes.Delete)
}

func main() {
	database.ConnectDb()
	models.Migrate()

	app := fiber.New()

	setUpRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":8080"))
}
