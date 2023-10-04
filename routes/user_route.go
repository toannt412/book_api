package routes

import (
	"bookstore/controllers"
	"bookstore/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("user/register", controllers.RegisterAccount())
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/user/:userId", controllers.GetUser())
			secured.PUT("/user/:userId", controllers.EditUser())
			secured.DELETE("/user/:userId", controllers.DeleteUser())
			secured.GET("/users", controllers.GetAllUsers())
		}
	}
	//router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetUser())
	router.PUT("/user/:userId", controllers.EditUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
	router.GET("/users", controllers.GetAllUsers())
	router.POST("/register", controllers.RegisterAccount())
	router.POST("/login", controllers.LoginAccount())
}
