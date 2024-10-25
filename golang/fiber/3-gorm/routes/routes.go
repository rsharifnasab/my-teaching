package routes

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm-postgres/models"
	"gorm-postgres/repository"
)

type HTTPServer struct {
	Repo repository.Book
}

func (server *HTTPServer) Register(app fiber.Router) {
	app.Get("/hello", server.Hello)
	app.Post("/addbook", server.AddBook)
	app.Get("/allbooks", server.AllBooks)
	app.Get("/books/:id/salam", server.GetBook)
}

// Hello
func (server *HTTPServer) Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

func (server *HTTPServer) GetBook(c *fiber.Ctx) error {
	bookID := c.Params("id")
	validIDs := []string{"42", "32", "22", "12"}
	for _, id := range validIDs {
		if id == bookID {
			return c.SendString(bookID)
		}
	}

	return c.SendStatus(404)
}

// AddBook
func (server *HTTPServer) AddBook(c *fiber.Ctx) error {
	println(c.Query("key1", "default"))
	book := &AddBookRequest{}
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	fmt.Printf("added book: %+v\n", book)

	newBook := models.Book{
		Title:  book.Title,
		Author: book.Author,
	}

	createdBook, err := server.Repo.Add(newBook)
	if err != nil {
		return err
	}

	result := AddBookResponse{
		Code:   http.StatusCreated,
		Status: "success",
		ID:     int(createdBook.ID),
	}

	return c.Status(http.StatusCreated).JSON(result)
}

// AllBooks
func (server *HTTPServer) AllBooks(c *fiber.Ctx) error {
	books, err := server.Repo.GetAll()
	if err != nil {
		return err
	}

	return c.Status(200).JSON(books)
}
