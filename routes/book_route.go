package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine) {
	//All routes related to books comes here

	router.POST("/book", controllers.CreateBook())
	router.GET("/book/:bookId", controllers.GetABook())
	router.PUT("/book/:bookId", controllers.EditABook())
	router.DELETE("/book/:bookId", controllers.DeleteABook())
	router.GET("/books", controllers.GetAllBooks())
}
