package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *Routes) CartRoute(router *gin.Engine) {
	router.POST("/cart", r.cartRoutes.CreateCart())
	router.GET("/cart/:cartId", r.cartRoutes.GetCart())
	router.PUT("/cart/:cartId", r.cartRoutes.EditACart())
	router.DELETE("/cart/:cartId", r.cartRoutes.DeleteCart())
}
