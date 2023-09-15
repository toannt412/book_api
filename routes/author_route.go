package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func AuthorRoute(router *gin.Engine) {

	router.POST("/author", controllers.CreateAuthor())
	router.GET("/authors", controllers.GetAllAuthors())
	router.PUT("/author/:authorId", controllers.EditAAuthor())
	router.DELETE("/author/:authorId", controllers.DeleteAAuthor())
	router.GET("/author/:authorId", controllers.GetAAuthor())
}
