package routes

import (
	"bookstore/controllers"
	"bookstore/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {
	order := router.Group("/order").Use(middlewares.AuthUser())
	{
		order.POST("/", controllers.CreateOrder())
		order.GET("/:orderId", controllers.GetOrder())
		order.PUT("/:orderId", controllers.EditOrder())
		order.DELETE("/:orderId", controllers.DeleteOrder())
	}
}
