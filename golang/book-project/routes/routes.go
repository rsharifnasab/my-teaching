package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/utils"
	"gorm-postgres/models"
	"gorm-postgres/repository"
)

type HTTPServer struct {
	Repo repository.Book
}

/*

// Create a new fiber instance with custom config
app := fiber.New(fiber.Config{
    // Override default error handler
    ErrorHandler: func(ctx *fiber.Ctx, err error) error {
        // Status code defaults to 500
        code := fiber.StatusInternalServerError

        // Retrieve the custom status code if it's a *fiber.Error
        var e *fiber.Error
        if errors.As(err, &e) {
            code = e.Code
        }

        // Send custom error page
        err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
        if err != nil {
            // In case the SendFile fails
            return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
        }

        // Return from handler
        return nil
    },
})

*/

func (server *HTTPServer) ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	return c.Status(code).JSON(fiber.Map{
		"status": "failed",
		"code":   code,
		"reason": err.Error(),
	})
}

func (server *HTTPServer) Register(app fiber.Router) {
	app.Get("/", server.Hello)
	app.Post("/books", server.AddBook)
	app.Post("/books-async", server.AddBookAsync)
	app.Get("/books", server.AllBooks)

	app.Get("/books/error", server.ErrorCreator)
	app.Get("/books/panic", server.Panicer)
	app.Get("/books/:id", server.GetBookByID)
}

func (server *HTTPServer) Hello(c *fiber.Ctx) error {
	return c.SendString("Hello from fiber")
}

func (server *HTTPServer) Panicer(c *fiber.Ctx) error {
	panic("sorry I panicked")
}

func (server *HTTPServer) ErrorCreator(c *fiber.Ctx) error {
	return errors.New("cannot process")
}

// TODO
func (server *HTTPServer) GetBookByID(c *fiber.Ctx) error {
	bookID := c.Params("id")

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		return c.SendStatus(400)
	}

	bookEntity, err := server.Repo.Get(uint64(bookIDInt))
	if err != nil {
		return c.SendStatus(404)
	}

	bookResponse := AddBookResponse{
		Title:  bookEntity.Title,
		Author: bookEntity.Author,
		ID:     int(bookEntity.ID),
	}

	return c.JSON(bookResponse)
}

func (server *HTTPServer) AddBookAsync(c *fiber.Ctx) error {
	bodyBytes := utils.CopyBytes(c.Body())
	book := &AddBookRequest{}
	if err := json.Unmarshal(bodyBytes, book); err != nil {
		return c.SendStatus(400)
	}

	go func(book *AddBookRequest) {
		newBook := models.Book{
			Title:  book.Title,
			Author: book.Author,
		}

		_, err := server.Repo.Add(newBook)
		if err != nil {
			log.Error("async add book encountered error: %s", err.Error())
			return
		}
	}(book)

	return c.SendString("ok")
}

func (server *HTTPServer) AddBook(c *fiber.Ctx) error {
	book := &AddBookRequest{}
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	newBook := models.Book{
		Title:  book.Title,
		Author: book.Author,
	}

	createdBook, err := server.Repo.Add(newBook)
	if err != nil {
		return err // problematic
	}

	result := AddBookResponse{
		Author: createdBook.Author,
		Title:  createdBook.Title,
		ID:     int(createdBook.ID),
	}

	location := fmt.Sprintf("%s:%d/v1/%v",
		"localhost",
		8080,
		createdBook.ID,
	)

	c.Response().Header.Add("Location", location)

	return c.Status(http.StatusCreated).JSON(result)
}

// TODO:: add pagniation
func (server *HTTPServer) AllBooks(c *fiber.Ctx) error {
	// sort,filter, page, page-size in query Param
	// + good default
	books, err := server.Repo.GetAll()
	if err != nil {
		return err
	}

	return c.Status(200).JSON(books)
}
