package controllers

import (
	"bookstore/responses"
	"bookstore/serialize"
	service "bookstore/service/category"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryController struct {
	categorySvc *service.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categorySvc: service.NewCategoryService(),
	}
}

// Create
func (ctrl *CategoryController) CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category *serialize.Category

		//validate the request body
		if err := c.BindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, responses.CategoryResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newCategory := &serialize.Category{
			Id:      primitive.NewObjectID(),
			CatName: category.CatName,
		}

		res, err := ctrl.categorySvc.CreateCategory(c, newCategory)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CategoryResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.CategoryResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

func (ctrl *CategoryController) GetCategoryByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryID := c.Param("categoryId")
		res, err := ctrl.categorySvc.GetCategoryByID(c, categoryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CategoryResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			responses.CategoryResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
		)
	}
}

// Read
// GET ALL
func (ctrl *CategoryController) GetAllCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := ctrl.categorySvc.GetAllCategories(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CategoryResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			responses.CategoryResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
		)
	}
}

// Update
func (ctrl *CategoryController) EditCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryId := c.Param("categoryId")
		var category *serialize.Category

		if err := c.BindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, responses.CategoryResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		update := &serialize.Category{
			CatName: category.CatName,
		}
		res, err := ctrl.categorySvc.EditCategory(c, categoryId, update)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CategoryResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.CategoryResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})

	}
}

// Delete
func (ctrl *CategoryController) DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryId := c.Param("categoryId")

		res, err := ctrl.categorySvc.DeleteCategory(c, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CategoryResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			responses.CategoryResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
		)
	}
}
