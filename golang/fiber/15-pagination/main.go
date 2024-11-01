package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Sample data: a list of strings
var items = []string{
	"Item 1", "Item 2", "Item 3", "Item 4", "Item 5",
	"Item 6", "Item 7", "Item 8", "Item 9", "Item 10",
	"Item 11", "Item 12", "Item 13", "Item 14", "Item 15",
	"Item 16", "Item 17", "Item 18", "Item 19", "Item 20",
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/items", paginatedItems)
	app.Get("/items-cursor", cursorBasedItems)

	log.Fatal(app.Listen(":3000"))
}

func paginatedItems(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "5"))
	if err != nil || limit < 1 {
		limit = 5
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(items) {
		return c.JSON(fiber.Map{
			"page":  page,
			"limit": limit,
			"items": []string{},
			"error": "No items on this page",
		})
	}

	if end > len(items) {
		end = len(items)
	}

	paginatedItems := items[start:end]

	return c.JSON(fiber.Map{
		"page":  page,
		"limit": limit,
		"items": paginatedItems,
	})
}

// GET http://localhost:3000/items?page=1&limit=5

func cursorBasedItems(c *fiber.Ctx) error {
	cursor, err := strconv.Atoi(c.Query("cursor", "0"))
	if err != nil || cursor < 0 {
		cursor = 0
	}

	limit, err := strconv.Atoi(c.Query("limit", "5"))
	if err != nil || limit < 1 {
		limit = 5
	}

	if cursor >= len(items) {
		return c.JSON(fiber.Map{
			"cursor": cursor,
			"limit":  limit,
			"items":  []string{},
			"error":  "No more items",
		})
	}

	end := cursor + limit
	if end > len(items) {
		end = len(items)
	}

	paginatedItems := items[cursor:end]

	var nextCursor *int
	if end < len(items) {
		nextCursorVal := end
		nextCursor = &nextCursorVal
	}

	return c.JSON(fiber.Map{
		"cursor":     cursor,
		"limit":      limit,
		"items":      paginatedItems,
		"nextCursor": nextCursor, // This will be nil if there are no more items
	})
}
