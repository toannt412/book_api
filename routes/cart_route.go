package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.Engine) {
	router.POST("/cart", controllers.CreateCart())
	router.GET("/cart/:cartId", controllers.GetACart())
	router.PUT("/cart/:cartId", controllers.EditACart())
	router.DELETE("/cart/:cartId", controllers.DeleteACart())
}
