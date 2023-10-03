package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func AuthorRoute(router *gin.Engine) {

	router.POST("/author", controllers.CreateAuthor())
	router.GET("/authors", controllers.GetAllAuthors())
	router.PUT("/author/:authorId", controllers.EditAuthor())
	router.DELETE("/author/:authorId", controllers.DeleteAuthor())
	router.GET("/author/:authorId", controllers.GetAuthor())
}
