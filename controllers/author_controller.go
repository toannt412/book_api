package controllers

import (
	"bookstore/responses"
	"bookstore/serialize"
	service "bookstore/service/author"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorController struct {
	authorSvc *service.AuthorService
}

func NewAuthorController() *AuthorController {
	return &AuthorController{
		authorSvc: service.NewAuthorService(),
	}
}
func (ctrl *AuthorController) CreateAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var author *serialize.Author

		//validate the request body
		if err := c.BindJSON(&author); err != nil {
			c.JSON(http.StatusBadRequest, responses.AuthorResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		// TIM HIEU NAY LAI
		// //use the validator library to validate required fields
		// if validationErr := validate.Struct(&user); validationErr != nil {
		// 	c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		// 	return
		// }

		newAuthor := &serialize.Author{
			Id:          primitive.NewObjectID(),
			AuthorName:  author.AuthorName,
			DateOfBirth: author.DateOfBirth,
			HomeTown:    author.HomeTown,
			Alive:       author.Alive,
		}

		res, err := ctrl.authorSvc.CreateAuthor(c, newAuthor)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AuthorResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AuthorResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

// Read
// GET ALL
func (ctrl *AuthorController) GetAllAuthors() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := ctrl.authorSvc.GetAllAuthors(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AuthorResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			responses.AuthorResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
		)
	}
}

// Update
func (ctrl *AuthorController) EditAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorId := c.Param("authorId")
		ojbId, _ := primitive.ObjectIDFromHex(authorId)
		var author *serialize.Author

		if err := c.BindJSON(&author); err != nil {
			c.JSON(http.StatusBadRequest, responses.AuthorResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		update := &serialize.Author{
			Id:          ojbId,
			AuthorName:  author.AuthorName,
			DateOfBirth: author.DateOfBirth,
			HomeTown:    author.HomeTown,
			Alive:       author.Alive,
		}

		res, err := ctrl.authorSvc.EditAuthor(c, authorId, update)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.AuthorResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.AuthorResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

// Delete
func (ctrl *AuthorController) DeleteAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorId := c.Param("authorId")

		res, err := ctrl.authorSvc.DeleteAuthor(c, authorId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AuthorResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			responses.AuthorResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
		)
	}
}

func (ctrl *AuthorController) GetAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorId := c.Param("authorId")

		res, err := ctrl.authorSvc.GetAuthorByID(c, authorId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AuthorResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AuthorResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}
