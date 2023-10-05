package routes

import (
	"bookstore/controllers"
	"bookstore/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {

	router.POST("/admin/login", controllers.LoginAccountAdmin())
	auth := router.Group("/auth").Use(middlewares.Auth())
	{
		auth.GET("/admins", controllers.GetAllAdmins())
		auth.GET("/admin/:adminId", controllers.GetAdmin())
		auth.PUT("/admin/:adminId", controllers.EditAdmin())
		auth.DELETE("/admin/:adminId", controllers.DeleteAdmin())
		auth.POST("/admin", controllers.CreateAdmin())
		auth.POST("/logout", middlewares.Logout())
	}

}
