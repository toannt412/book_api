package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *Routes) BookRoute(router *gin.Engine) {
	//All routes related to books comes here

	router.POST("/book", r.bookRoutes.CreateBook())
	router.GET("/book/:bookId", r.bookRoutes.GetBook())
	router.PUT("/book/:bookId", r.bookRoutes.EditBook())
	router.DELETE("/book/:bookId", r.bookRoutes.DeleteBook())
	router.GET("/books", r.bookRoutes.GetAllBooks())
}
