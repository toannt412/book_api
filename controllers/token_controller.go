package controllers

import (
	"bookstore/auth"
	"bookstore/helpers"
	service "bookstore/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	//var user model.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	// check if email exists and password is correct
	record, err := service.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	credentialError := helpers.CheckPasswordHash(record.Password, request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": credentialError.Error()})
		c.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(record.Email, record.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
