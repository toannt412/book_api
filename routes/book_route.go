package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine) {
	//All routes related to books comes here

	router.POST("/book", controllers.CreateBook())
	router.GET("/book/:bookId", controllers.GetBook())
	router.PUT("/book/:bookId", controllers.EditBook())
	router.DELETE("/book/:bookId", controllers.DeleteBook())
	router.GET("/books", controllers.GetAllBooks())
}
