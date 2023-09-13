package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoute(router *gin.Engine) {

	router.POST("/register", controllers.RegisterAccount())
	router.POST("/login", controllers.LoginAccount())
	router.GET("/logout", CheckToken(), controllers.LogoutAccount())

}
