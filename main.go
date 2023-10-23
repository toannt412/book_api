package main

import (
	//"net/http"

	"bookstore/configs"
	"bookstore/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	routes *routes.Routes
}

func NewRoutes(router *gin.Engine) *Routes {
	return &Routes{
		routes: routes.NewRoutes(router),
	}
}
func main() {
	configs.Load()
	router := gin.Default()
	//run database
	//dao.ConnectDB()
	r := routes.NewRoutes(router)
	//routes
	// routes.UserRoute(router)
	// routes.BookRoute(router)
	// routes.CategoryRoute(router)
	// routes.AuthorRoute(router)
	r.AdminRoute(router)
	r.AuthorRoute(router)
	r.BookRoute(router)
	r.CartRoute(router)
	r.CategoryRoute(router)
	r.OrderRoute(router)
	r.UserRoute(router)
	// routes.OrderRoute(router)
	// routes.CartRoute(router)
	//routes.NewRoutes(router)
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"data": "Hello from Gin-gonic & MongoDB",
	// 	})
	// })
	router.Run(fmt.Sprintf(":%s", configs.Config.Port))
}
