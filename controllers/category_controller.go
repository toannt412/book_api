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

var categoriesCollection *mongo.Collection = configs.GetCollection(configs.DB, "categories")

// Create
func CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var category models.Category
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, responses.CategoryResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		// TIM HIEU NAY LAI
		// //use the validator library to validate required fields
		// if validationErr := validate.Struct(&user); validationErr != nil {
		// 	c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		// 	return
		// }

		newCategory := models.Category{
			Id:      primitive.NewObjectID(),
			CatName: category.CatName,
		}

		result, err := categoriesCollection.InsertOne(ctx, newCategory)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CategoryResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.CategoryResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

// Read
// GET ALL
func GetAllCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var categories []models.Category
		defer cancel()

		results, err := categoriesCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CategoryResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleCategory models.Category
			if err = results.Decode(&singleCategory); err != nil {
				c.JSON(http.StatusInternalServerError, responses.CategoryResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			categories = append(categories, singleCategory)
		}

		c.JSON(http.StatusOK,
			responses.CategoryResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": categories}},
		)
	}
}

// Update
func EditACategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		categoryId := c.Param("categoryId")
		var category models.Category
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(categoryId)

		if err := c.BindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, responses.CategoryResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		update := models.Category{
			Id:      objId,
			CatName: category.CatName,
		}

		result, err := categoriesCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CategoryResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated book details
		var updatedCategory models.Category
		if result.MatchedCount == 1 {
			err := categoriesCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedCategory)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.CategoryResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

		}

		c.JSON(http.StatusOK, responses.CategoryResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedCategory}})

	}
}

// Delete
func DeleteACategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		categoryId := c.Param("categoryId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(categoryId)

		result, err := categoriesCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CategoryResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.CategoryResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Category with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.CategoryResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Category successfully deleted!"}},
		)
	}
}
