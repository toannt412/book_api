package routes

import (
	"bookstore/middlewares"

	"github.com/gin-gonic/gin"
)

func (r *Routes) OrderRoute(router *gin.Engine) {
	order := router.Group("/order").Use(middlewares.NewMiddlewares().AuthUser())
	{
		order.POST("/", r.orderRoutes.CreateOrder())
		order.GET("/:orderId", r.orderRoutes.GetOrder())
		order.PUT("/:orderId", r.orderRoutes.EditOrder())
		order.DELETE("/:orderId", r.orderRoutes.DeleteOrder())
	}
}
