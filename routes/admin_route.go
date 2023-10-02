package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {
	router.POST("/admin/login", controllers.LoginAccountAdmin())
	router.GET("/admins", controllers.GetAllAdmins())
	router.GET("/admin/:adminId", controllers.GetAdmin())
	router.PUT("/admin/:adminId", controllers.EditAdmin())
	router.DELETE("/admin/:adminId", controllers.DeleteAdmin())
	router.POST("/admin", controllers.CreateAdmin())
}
