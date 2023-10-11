package routes

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	adminRoutes    *controllers.AdminController
	authorRoutes   *controllers.AuthorController
	bookRoutes     *controllers.BookController
	cartRoutes     *controllers.CartController
	categoryRoutes *controllers.CategoryController
	orderRoutes    *controllers.OrderController
	userRoutes     *controllers.UserController
}

func NewRoutes(router *gin.Engine) *Routes {
	return &Routes{
		adminRoutes:    controllers.NewAdminController(),
		authorRoutes:   controllers.NewAuthorController(),
		bookRoutes:     controllers.NewBookController(),
		cartRoutes:     controllers.NewCartController(),
		categoryRoutes: controllers.NewCategoryController(),
		orderRoutes:    controllers.NewOrderController(),
		userRoutes:     controllers.NewUserController(),
	}
}
