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
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

var authorsCollection *mongo.Collection = configs.GetCollection(configs.DB, "authors")

func CreateAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var author models.Author
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&author); err != nil {
			c.JSON(http.StatusBadRequest, responses.AuthorResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		// TIM HIEU NAY LAI
		// //use the validator library to validate required fields
		// if validationErr := validate.Struct(&user); validationErr != nil {
		// 	c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		// 	return
		// }

		newAuthor := models.Author{
			Id:          primitive.NewObjectID(),
			AuthorName:  author.AuthorName,
			DateOfBirth: author.DateOfBirth,
			HomeTown:    author.HomeTown,
			Alive:       author.Alive,
		}

		result, err := authorsCollection.InsertOne(ctx, newAuthor)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AuthorResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AuthorResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

// Read
// GET ALL
func GetAllAuthors() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var authors []models.Author
		defer cancel()

		results, err := authorsCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AuthorResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleAuthor models.Author
			if err = results.Decode(&singleAuthor); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AuthorResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			authors = append(authors, singleAuthor)
		}

		c.JSON(http.StatusOK,
			responses.AuthorResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": authors}},
		)
	}
}

// Update
func EditAAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		authorId := c.Param("authorId")
		var author models.Author
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(authorId)
		opts := options.FindOneAndUpdate().SetUpsert(false)

		if err := c.BindJSON(&author); err != nil {
			c.JSON(http.StatusBadRequest, responses.AuthorResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		update := models.Author{
			Id:          objId,
			AuthorName:  author.AuthorName,
			DateOfBirth: author.DateOfBirth,
			HomeTown:    author.HomeTown,
			Alive:       author.Alive,
		}

		err := authorsCollection.FindOneAndUpdate(ctx, bson.M{"_id": objId}, bson.M{"$set": update}, opts).Decode(&author)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AuthorResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated book details
		var updatedAuthor models.Author
		if err := authorsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedAuthor); err != nil {

			c.JSON(http.StatusInternalServerError, responses.AuthorResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return

		}
		c.JSON(http.StatusOK, responses.AuthorResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedAuthor}})
	}
}

// Delete
func DeleteAAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		authorId := c.Param("authorId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(authorId)

		result, err := authorsCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AuthorResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.AuthorResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Author with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.AuthorResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Author successfully deleted!"}},
		)
	}
}

func GetAAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		authorId := c.Param("authorId")
		var author models.Author
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(authorId)

		err := authorsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&author)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AuthorResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AuthorResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": author}})
	}
}
