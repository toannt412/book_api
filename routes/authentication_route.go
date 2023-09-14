package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoute(router *gin.Engine) {



	router.GET("/logout", CheckToken(), controllers.LogoutAccount())
	router.PUT("/auth/:accountId", controllers.EditAUserAccount())
	router.GET("/auth/:accountId", controllers.GetInfoAUserAuth())

}
