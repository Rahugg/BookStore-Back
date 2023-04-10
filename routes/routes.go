package routes

import (
	"Assignment3Go/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("api/book", controllers.CreateBook)
	app.Get("api/books", controllers.GetBooks)
	app.Patch("api/book/:id", controllers.UpdateBook)
	app.Delete("api/book/:id", controllers.DeleteBook)
	app.Get("api/book", controllers.SearchBook)
	app.Get("api/sortedBooksAscending", controllers.SortedBooksASC)
	app.Get("api/sortedBooksDescending", controllers.SortedBooksDESC)
}
