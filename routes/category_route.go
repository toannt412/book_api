package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {
	router.POST("/category", controllers.CreateCategory())
	router.GET("/category/:categoryId", controllers.GetCategoryByID())
	router.GET("/categories", controllers.GetAllCategories())
	router.PUT("/category/:categoryId", controllers.EditCategory())
	router.DELETE("/category/:categoryId", controllers.DeleteCategory())
}
