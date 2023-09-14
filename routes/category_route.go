package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {
	router.POST("/category", controllers.CreateCategory())
	router.GET("/categories", controllers.GetAllCategories())
	router.PUT("/category/:categoryId", controllers.EditACategory())
	router.DELETE("/category/:categoryId", controllers.DeleteACategory())
}
