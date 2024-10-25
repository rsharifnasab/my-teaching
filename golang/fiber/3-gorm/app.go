package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm-postgres/database"
	"gorm-postgres/models"
	"gorm-postgres/repository"
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
	db := database.ConnectDb()
	bookRepo := repository.NewGormBook(db)
	_ = bookRepo

	models.Migrate()

	app := fiber.New()

	app.Get("/allbooks", routes.AllBooks)
	app.Get("/allbooks2", routes.AllBooksCreator(bookRepo))

	setUpRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":8080"))
}
