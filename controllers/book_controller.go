package controllers

import (
	"bookstore/responses"
	"bookstore/serialize"
	service "bookstore/service/book"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookController struct {
	bookSvc *service.BookService
}

func NewBookController() *BookController {
	return &BookController{
		bookSvc: service.NewBookService(),
	}
}

// Create
func (ctrl *BookController) CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book *serialize.Book
		//validate the request body
		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		newBook := &serialize.Book{
			Id:                primitive.NewObjectID(),
			BookName:          book.BookName,
			Price:             book.Price,
			PublishingCompany: book.PublishingCompany,
			PublicationDate:   book.PublicationDate,
			Description:       book.Description,
			CategoryIDs:       book.CategoryIDs,
			AuthorID:          book.AuthorID,
		}

		res, err := ctrl.bookSvc.CreateBook(c, newBook)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.BookResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"data": res}})
	}
}

// Read
// GET BY ID
func (ctrl *BookController) GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookId := c.Param("bookId")

		res, err := ctrl.bookSvc.GetBookByID(c, bookId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.BookResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": res}})
	}
}

// GET ALL
func (ctrl *BookController) GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := ctrl.bookSvc.GetAllBooks(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
		)
	}
}

// Update
func (ctrl *BookController) EditBook() gin.HandlerFunc {
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
		res, err := ctrl.bookSvc.EditBook(c, bookId, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})

	}
}

// Delete
func (ctrl *BookController) DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {

		bookId := c.Param("bookId")

		res, err := ctrl.bookSvc.DeleteBook(c, bookId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return

		}
		c.JSON(http.StatusOK,
			responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
		)
	}
}
