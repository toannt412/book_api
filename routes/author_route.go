package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *Routes) AuthorRoute(router *gin.Engine) {

	router.POST("/author", r.authorRoutes.CreateAuthor())
	router.GET("/authors", r.authorRoutes.GetAllAuthors())
	router.PUT("/author/:authorId", r.authorRoutes.EditAuthor())
	router.DELETE("/author/:authorId", r.authorRoutes.DeleteAuthor())
	router.GET("/author/:authorId", r.authorRoutes.GetAuthor())
}
