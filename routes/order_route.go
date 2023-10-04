package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {
	router.POST("/order", controllers.CreateOrder())
	router.GET("/order/:orderId", controllers.GetOrder())
	router.PUT("/order/:orderId", controllers.EditOrder())
	router.DELETE("/order/:orderId", controllers.DeleteOrder())
}
