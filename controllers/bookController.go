package controllers

import (
	"Assignment3Go/database"
	"Assignment3Go/models"
	"github.com/gofiber/fiber/v2"
)

func CreateBook(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	book := models.Book{
		Title:       data["title"],
		Description: data["description"],
		Cost:        data["cost"],
	}

	database.DB.Create(&book)

	return c.JSON(book)
}

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Find(&books)

	c.JSON(books)
	//fmt.Println(books)
	return c.SendStatus(200)
}

func UpdateBook(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}
	database.DB.Model(models.Book{}).Where("id = ?", data["id"]).Updates(models.Book{Title: data["title"]})
	database.DB.Model(models.Book{}).Where("id = ?", data["id"]).Updates(models.Book{Description: data["description"]})
	return c.SendStatus(200)
}

func DeleteBook(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}
	var book models.Book
	var deletingId = data["id"]
	database.DB.Delete(&book, deletingId)

	//db.Delete(&User{}, 10)
	return c.SendStatus(200)
}

func SearchBook(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	var book models.Book

	database.DB.Where("name=?", data["name"]).First(&book)

	if book.Title == "" {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "item not found",
		})
	}

	return c.JSON(book)
}

func SortedBooksASC(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Order("cost").Find(&books)

	c.JSON(books)
	//fmt.Println(books)
	return c.SendStatus(200)
}
func SortedBooksDESC(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Order("cost desc").Find(&books)

	c.JSON(books)
	//fmt.Println(books)
	return c.SendStatus(200)
}
