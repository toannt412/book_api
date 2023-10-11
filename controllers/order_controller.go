package controllers

import (
	"bookstore/responses"
	"bookstore/serialize"
	service "bookstore/service/order"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController struct {
	orderSvc *service.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		orderSvc: service.NewOrderService(),
	}
}
func (ctrl *OrderController) CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order *serialize.Order

		//validate the request body
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		newOrder := &serialize.Order{
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

		res, err := ctrl.orderSvc.CreateOrder(c, newOrder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.OrderResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

func (ctrl *OrderController) GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("orderId")

		res, err := ctrl.orderSvc.GetOrderByID(c, orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.OrderResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

func (ctrl *OrderController) DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("orderId")
		res, err := ctrl.orderSvc.DeleteOrder(c, orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.OrderResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

func (ctrl *OrderController) EditOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("orderId")
		var order *serialize.Order
		objId, _ := primitive.ObjectIDFromHex(orderId)

		//validate the request body
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		updateOrder := &serialize.Order{
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

		res, err := ctrl.orderSvc.EditOrder(c, orderId, updateOrder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.OrderResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}
