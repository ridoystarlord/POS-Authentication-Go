package controllers

import (
	"authentication/models"
	"authentication/storage"

	"github.com/gofiber/fiber/v2"
)

//  Books godoc
//	@Summary		Create book
//	@Description	Create a new book
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			request	body	models.Book	true	"Book"
//	@Success		200
//	@failure		400	{string}	string	"error"
//	@Router			/book/new [post]
//	@Security		BearerAuth
func  CreateBook(c *fiber.Ctx) error {
	var book models.Book
	err := c.BodyParser(&book)
	if  err != nil {
		 c.JSON(&fiber.Map{"message":"Unable to parse body"})
		 return err
	}
	err=storage.DB.Create(&book).Error
	if err != nil {
		c.JSON(&fiber.Map{"message":"Unable to create book"})
		 return err
	}
	c.JSON(&fiber.Map{"message":"Book created successfully"})
	return nil
}


//  Books godoc
//	@Summary		Delete book by id
//	@Description	Delete a single book from the database
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Book ID"
//	@Success		200	
//	@Router			/book/{id} [delete]
//	@Security		BearerAuth
func  DeleteBook(c *fiber.Ctx) error {
	var book models.Book
	id := c.Params("id")
	err := storage.DB.Delete(&book,id).Error
	if err != nil {
		c.JSON(&fiber.Map{"message":"Unable to delete book"})
		 return err
	}
	c.JSON(&fiber.Map{"message":"Book deleted successfully"})
	return nil
}

//  Books godoc
//	@Summary		List of books
//	@Description	Get all the books from the database
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.Book
//	@Router			/book [get]
//	@Security		BearerAuth
func  GetBook(c *fiber.Ctx) error {
	var books []models.Book
	err := storage.DB.Find(&books).Error
	if err != nil {
		c.JSON(&fiber.Map{"message":"Unable to get books"})
		 return err
	}
	c.JSON(&books)
	return nil
}


//  Books godoc
//	@Summary		Get book by id
//	@Description	Read a single book from the database
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		200	{object}	models.Book
//	@Router			/book/{id} [get]
//	@Security		BearerAuth
func  GetBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	err := storage.DB.First(&book, id).Error
	if err != nil {
		c.JSON(&fiber.Map{"message":"Unable to get book"})
		 return err
	}
	c.JSON(&book)
	return nil
}