package routes

import (
	"bookstore/middlewares"

	"github.com/gin-gonic/gin"
)

func (r *Routes) UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	api := router.Group("/api")
	{
		//api.POST("/token", controllers.GenerateToken)
		api.POST("user/register", r.userRoutes.RegisterAccount())
		api.POST("user/login", r.userRoutes.LoginAccount())
		user := api.Group("/user").Use(middlewares.NewMiddlewares().AuthUser())
		{
			user.GET("/:userId", r.userRoutes.GetUser())
			user.PUT("/:userId", r.userRoutes.EditUser())
			user.DELETE("/:userId", r.userRoutes.DeleteUser())
			user.POST("/logout", middlewares.NewMiddlewares().LogoutUser())
			//secured.GET("/users", r.userRoutes.GetAllUsers())
		}
		api.POST("/forgot-password", r.userRoutes.ForgotPassword())
		api.POST("/reset-password", r.userRoutes.ResetPassword())
		//forgotPass := api.Group("/forgot-password").Use(middlewares.CheckEmail())
		// {
		// 	forgotPass.POST("/reset", r.userRoutes.ResetPassword())
		// }
	}

}
