package controllers

import (
	"bookstore/dao/book/model"
	"bookstore/responses"
	"bookstore/serialize"
	service "bookstore/service/book"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create
func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book model.Book
		//validate the request body
		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		// TIM HIEU NAY LAI
		// //use the validator library to validate required fields
		// if validationErr := validate.Struct(&user); validationErr != nil {
		// 	c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		// 	return
		// }

		newBook := model.Book{
			Id:                primitive.NewObjectID(),
			BookName:          book.BookName,
			Price:             book.Price,
			PublishingCompany: book.PublishingCompany,
			PublicationDate:   book.PublicationDate,
			Description:       book.Description,
			CategoryIDs:       book.CategoryIDs,
			AuthorID:          book.AuthorID,
		}

		res, err := service.CreateBook(c, newBook)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.BookResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

// Read
// GET BY ID
func GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookId := c.Param("bookId")

		res, err := service.GetBookByID(c, bookId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

// Update
func EditBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookId := c.Param("bookId")
		objID, _ := primitive.ObjectIDFromHex(bookId)
		var book *serialize.Book

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		update := &serialize.Book{
			Id:                objID,
			BookName:          book.BookName,
			Price:             book.Price,
			PublishingCompany: book.PublishingCompany,
			PublicationDate:   book.PublicationDate,
			Description:       book.Description,
			CategoryIDs:       book.CategoryIDs,
			AuthorID:          book.AuthorID,
		}
		res, err := service.EditBook(c, bookId, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})

	}
}

// Delete
func DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {

		bookId := c.Param("bookId")

		res, err := service.DeleteBook(c, bookId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return

		}
		c.JSON(http.StatusOK,
			responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
		)
	}
}

// GET ALL
func GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {

		res, err := service.GetAllBooks(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
		)
	}
}
