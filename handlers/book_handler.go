package handlers

import (
	"strconv"

	"github.com/dekindrama/go-book-api/databases"
	"github.com/dekindrama/go-book-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (handler *BookHandler) GetBooks(context *fiber.Ctx) error {
	//* get all books
	var books []models.BookModel
	if err := databases.Mysql.Find(&books).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return context.Status(fiber.StatusOK).JSON(&books)
}

func (handler *BookHandler) GetBook(context *fiber.Ctx) error {
	//* find book
	id, _ := strconv.Atoi(context.Params("id"))
	var book models.BookModel
	if err := databases.Mysql.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "book not found"})
		}
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return context.Status(fiber.StatusOK).JSON(&book)
}

func (handler *BookHandler) StoreBook(context *fiber.Ctx) error {
	//* parsing request
	var bookModel models.BookModel
	if err := context.BodyParser(&bookModel); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	//* store book
	if err := databases.Mysql.Create(&bookModel).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return context.Status(fiber.StatusCreated).JSON(&bookModel)
}

func (handler *BookHandler) UpdateBook(context *fiber.Ctx) error {
	//* find book
	id, _ := strconv.Atoi(context.Params("id"))
	var book models.BookModel
	if err := databases.Mysql.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "book not found"})
		}
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	//* parsing request
	if err := context.BodyParser(&book); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	//* update book
	if err := databases.Mysql.Save(&book).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success update book"})
}

func (handler *BookHandler) DeleteBook(context *fiber.Ctx) error {
	//* delete book
	id, _ := strconv.Atoi(context.Params("id"))
	if err := databases.Mysql.Delete(&models.BookModel{}, id).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success delete book"})
}
