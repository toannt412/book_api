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

var cartCollection *mongo.Collection = configs.GetCollection(configs.DB, "carts")

func CreateCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		var cart models.Cart
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newCart := models.Cart{
			Id:            primitive.NewObjectID(),
			UserID:        cart.UserID,
			Books:         cart.Books,
			TotalQuantity: cart.TotalQuantity,
			TotalAmount:   cart.TotalAmount,
		}

		result, err := cartCollection.InsertOne(ctx, newCart)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.CartResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetACart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cartId := c.Param("cartId")
		var cart models.Cart
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(cartId)

		err := cartCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&cart)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": cart}})
	}
}

func DeleteACart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cartId := c.Param("cartId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(cartId)

		result, err := cartCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, responses.CartResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "cart with specified ID not found!"}})
		}
		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Cart successfully deleted!"}})
	}
}

func EditACart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cartId := c.Param("cartId")
		var cart models.Cart
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(cartId)

		//validate the request body
		if err := c.BindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		updatecart := models.Cart{
			Id:            objId,
			UserID:        cart.UserID,
			Books:         cart.Books,
			TotalQuantity: cart.TotalQuantity,
			TotalAmount:   cart.TotalAmount,
		}

		result, err := cartCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updatecart})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedcart models.Cart
		if result.MatchedCount == 1 {
			err := cartCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedcart)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatecart}})
	}
}
