package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *Routes) CategoryRoute(router *gin.Engine) {
	router.POST("/category", r.categoryRoutes.CreateCategory())
	router.GET("/category/:categoryId", r.categoryRoutes.GetCategoryByID())
	router.GET("/categories", r.categoryRoutes.GetAllCategories())
	router.PUT("/category/:categoryId", r.categoryRoutes.EditCategory())
	router.DELETE("/category/:categoryId", r.categoryRoutes.DeleteCategory())
}
