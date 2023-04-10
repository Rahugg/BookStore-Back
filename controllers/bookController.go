package controllers

import (
	"Assignment3Go/database"
	"Assignment3Go/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

type PatchData struct {
	Title       string `json:Title`
	Description string `json:Description`
}

func CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Create(&book)
	return c.SendStatus(200)

}

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Find(&books)

	c.JSON(books)
	//fmt.Println(books)
	return c.SendStatus(200)
}

func UpdateBook(c *fiber.Ctx) error {
	var patchData PatchData
	if err := c.BodyParser(&patchData); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Get the ID from the request parameters
	id := c.Params("id")

	// Find the book in the database
	var book models.Book
	if result := database.DB.First(&book, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "book not found",
			})
		}
		log.Println("Error retrieving book from database:", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error retrieving book from database",
		})
	}

	// Update the book fields with the patch data
	book.Title = patchData.Title
	book.Description = patchData.Description

	// Save the book in the database
	if result := database.DB.Save(&book); result.Error != nil {
		log.Println("Error updating book in database:", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating book in database",
		})
	}

	// Return a success response
	return c.JSON(fiber.Map{
		"message": "book updated successfully",
		"book":    book,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	// Query the database for the item with the specified ID
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		// Return an error if the item does not exist
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "book not found",
		})
	}

	// Delete the item from the database
	if err := database.DB.Delete(&book).Error; err != nil {
		// Return an error if the deletion fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting book",
		})
	}

	// Return a success message if the deletion succeeds
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("book %d deleted", book.ID),
	})
}

func SearchBook(c *fiber.Ctx) error {
	searchTerm := c.Query("searchTerm")

	var searchResults []models.Book
	database.DB.Where("title=?", searchTerm).Find(&searchResults)
	if len(searchResults) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "book not found",
		})
	}

	return c.JSON(searchResults)
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
