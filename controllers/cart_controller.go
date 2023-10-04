package controllers

import (
	"bookstore/responses"
	"bookstore/serialize"
	service "bookstore/service/cart"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cart *serialize.Cart

		//validate the request body
		if err := c.BindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newCart := &serialize.Cart{
			Id:            primitive.NewObjectID(),
			UserID:        cart.UserID,
			Books:         cart.Books,
			TotalQuantity: cart.TotalQuantity,
			TotalAmount:   cart.TotalAmount,
		}

		res, err := service.CreateCart(c, newCart)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.CartResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

func GetCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		cartId := c.Param("cartId")

		res, err := service.GetCart(c, cartId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

func DeleteCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		cartId := c.Param("cartId")

		res, err := service.DeleteCart(c, cartId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

func EditACart() gin.HandlerFunc {
	return func(c *gin.Context) {
		cartId := c.Param("cartId")
		var cart *serialize.Cart
		objId, _ := primitive.ObjectIDFromHex(cartId)

		//validate the request body
		if err := c.BindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		updatecart := &serialize.Cart{
			Id:            objId,
			UserID:        cart.UserID,
			Books:         cart.Books,
			TotalQuantity: cart.TotalQuantity,
			TotalAmount:   cart.TotalAmount,
		}

		res, err := service.EditCart(c, cartId, updatecart)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}
		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}
