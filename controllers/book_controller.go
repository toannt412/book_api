package controllers

import (
	"bookstore/configs"
	"bookstore/dao/models"
	"bookstore/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookCollection *mongo.Collection = configs.GetCollection(configs.DB, "books")

// Create
func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		var book models.Book
		defer cancel()

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

		newBook := models.Book{
			Id:                primitive.NewObjectID(),
			BookName:          book.BookName,
			Price:             book.Price,
			PublishingCompany: book.PublishingCompany,
			PublicationDate:   book.PublicationDate,
			Description:       book.Description,
			CategoryIDs:       book.CategoryIDs,
			AuthorID:          book.AuthorID,
		}

		result, err := bookCollection.InsertOne(ctx, newBook)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.BookResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

// Read
// GET BY ID
func GetABook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookId := c.Param("bookId")
		var book models.Book
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bookId)

		err := bookCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": book}})
	}
}

// Update
func EditABook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookId := c.Param("bookId")
		var book models.Book
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(bookId)

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// update := bson.M{"bookName": book.BookName,
		// 	"price":             book.Price,
		// 	"author":            book.Author,
		// 	"category":          book.Category,
		// 	"publishingCompany": book.PublishingCompany,
		// 	"publicationDate":   book.PublicationDate,
		// 	"description":       book.Description,
		// }
		update := models.Book{
			Id:                objId,
			BookName:          book.BookName,
			Price:             book.Price,
			PublishingCompany: book.PublishingCompany,
			PublicationDate:   book.PublicationDate,
			Description:       book.Description,
			CategoryIDs:       book.CategoryIDs,
			AuthorID:          book.AuthorID,
		}

		result, err := bookCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated book details
		var updatedBook models.Book
		if result.MatchedCount == 1 {
			err := bookCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedBook)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

		}

		c.JSON(http.StatusOK, responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedBook}})

	}
}

// Delete
func DeleteABook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookId := c.Param("bookId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bookId)

		result, err := bookCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.BookResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Book with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Book successfully deleted!"}},
		)
	}
}

// GET ALL
func GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var books []models.Book
		defer cancel()

		results, err := bookCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleBook models.Book
			if err = results.Decode(&singleBook); err != nil {
				c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			books = append(books, singleBook)
		}

		c.JSON(http.StatusOK,
			responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": books}},
		)
	}
}

// func ConvertStringToObjectId(data struct, str string){
// 	s := string()
// 	return json.Unmarshal([]byte(str), &data)
// }
