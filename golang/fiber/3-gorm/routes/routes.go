package routes

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm-postgres/database"
	"gorm-postgres/models"
)

// Hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

// AddBook
func AddBook(c *fiber.Ctx) error {
	book := new(AddBookRequest)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	fmt.Printf("added book: %+v\n", book)

	newBook := &models.Book{
		Title:  book.Title,
		Author: book.Author,
	}

	database.DB().Db.Create(newBook)

	result := AddBookResponse{
		Code:   http.StatusCreated,
		Status: "success",
		ID:     int(newBook.ID),
	}

	return c.Status(http.StatusCreated).JSON(result)
}

// AllBooks
func AllBooks(c *fiber.Ctx) error {
	books := []models.Book{}
	database.DB().Db.Find(&books)

	return c.Status(200).JSON(books)
}

// Book
func Book(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB().Db.Where("title = ?", title.Title).Find(&book)
	return c.Status(200).JSON(book)
}

// Update
func Update(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB().Db.Model(&book).Where("title = ?", title.Title).Update("author", title.Author)

	return c.Status(200).JSON("updated")
}

// Delete
func Delete(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB().Db.Where("title = ?", title.Title).Delete(&book)

	return c.Status(200).JSON("deleted")
}