package main

import (
	"log"

	"github.com/dekindrama/go-book-api/databases"
	"github.com/dekindrama/go-book-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//* init
	app := fiber.New()
	databases.NewMysql()

	//* routes
	BookHandler := handlers.NewBookHandler()
	routeBook := app.Group("/books")
	routeBook.Get("/", BookHandler.GetBooks)
	routeBook.Get("/:id", BookHandler.GetBook)
	routeBook.Post("/", BookHandler.StoreBook)
	routeBook.Put("/:id", BookHandler.UpdateBook)
	routeBook.Delete("/:id", BookHandler.DeleteBook)

	//* run project
	log.Fatal(app.Listen("127.0.0.1:8000"))
}
