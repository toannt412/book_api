package routes

import (
	"bookstore/middlewares"

	"github.com/gin-gonic/gin"
)

func (r *Routes) AdminRoute(router *gin.Engine) {

	router.POST("/admin/login", r.adminRoutes.LoginAccountAdmin())
	auth := router.Group("/auth").Use(middlewares.NewMiddlewares().AuthAdmin())
	{
		auth.GET("/admins", r.adminRoutes.GetAllAdmins())
		auth.GET("/admin/:adminId", r.adminRoutes.GetAdmin())
		auth.PUT("/admin/:adminId", r.adminRoutes.EditAdmin())
		auth.DELETE("/admin/:adminId", r.adminRoutes.DeleteAdmin())
		auth.POST("/admin", r.adminRoutes.CreateAdmin())
		auth.POST("/logout", middlewares.NewMiddlewares().LogoutAdmin())
	}

}
