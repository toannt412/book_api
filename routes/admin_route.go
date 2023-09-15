package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {
	router.POST("/admin/login", controllers.LoginAccountAdmin())
	router.GET("/admins", controllers.GetAllAdmins())
	router.GET("/admin/:adminId", controllers.GetAAdmin())
	router.PUT("/admin/:adminId", controllers.EditAAdmin())
	router.DELETE("/admin/:adminId", controllers.DeleteAAdmin())
	router.POST("/admin", controllers.CreateAdmin())
}
