package controllers

import (
	"bookstore/configs"
	"bookstore/models"
	"bookstore/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = configs.GetCollection(configs.DB, "orders")

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		var order models.Order
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newOrder := models.Order{
			Id:            primitive.NewObjectID(),
			UserID:        order.UserID,
			Books:         order.Books,
			CartID:        order.CartID,
			TotalQuantity: order.TotalQuantity,
			TotalPrice:    order.TotalPrice,
			TotalAmount:   order.TotalAmount,
			Status:        order.Status,
			OrderDate:     time.Now(),
		}

		result, err := orderCollection.InsertOne(ctx, newOrder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.OrderResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		orderId := c.Param("orderId")
		var order models.Order
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(orderId)

		err := orderCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.OrderResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": order}})
	}
}

func DeleteAOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		orderId := c.Param("orderId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(orderId)

		result, err := orderCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, responses.OrderResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Order with specified ID not found!"}})
		}
		c.JSON(http.StatusOK, responses.OrderResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Order successfully deleted!"}})
	}
}

func EditAOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		orderId := c.Param("orderId")
		var order models.Order
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(orderId)

		//validate the request body
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		updateOrder := models.Order{
			Id:            objId,
			UserID:        order.UserID,
			Books:         order.Books,
			CartID:        order.CartID,
			TotalQuantity: order.TotalQuantity,
			TotalPrice:    order.TotalPrice,
			TotalAmount:   order.TotalAmount,
			OrderDate:     time.Now(),
			Status:        order.Status,
		}

		result, err := orderCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updateOrder})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedOrder models.Order
		if result.MatchedCount == 1 {
			err := orderCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedOrder)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusOK, responses.OrderResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updateOrder}})
	}
}
