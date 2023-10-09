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
		//api.POST("/token", controllers.GenerateToken)
		api.POST("user/register", controllers.RegisterAccount())
		api.POST("user/login", controllers.LoginAccount())
		user := api.Group("/user").Use(middlewares.AuthUser())
		{
			user.GET("/:userId", controllers.GetUser())
			user.PUT("/:userId", controllers.EditUser())
			user.DELETE("/:userId", controllers.DeleteUser())
			user.POST("/logout", middlewares.LogoutUser())
			//secured.GET("/users", controllers.GetAllUsers())
		}
	}
	// //router.POST("/user", controllers.CreateUser())
	// router.GET("/user/:userId", controllers.GetUser())
	// router.PUT("/user/:userId", controllers.EditUser())
	// router.DELETE("/user/:userId", controllers.DeleteUser())
	// router.GET("/users", controllers.GetAllUsers())
	// router.POST("/register", controllers.RegisterAccount())
	// router.POST("/login", controllers.LoginAccount())
}
