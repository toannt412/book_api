package controllers

import (
	"bookstore/configs"
	"bookstore/helpers"
	"context"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var authenticationCollection *mongo.Collection = configs.GetCollection(configs.DB, "authentications")

func RegisterAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")

		if username == "" || password == "" || email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing username, password or email"})
			return
		}

		if govalidator.IsNull(username) || govalidator.IsNull(email) || govalidator.IsNull(password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User name can not empty"})
			return
		}

		if !govalidator.IsEmail(email) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is invalid"})
			return
		}

		username = helpers.Santize(username)
		password = helpers.Santize(password)
		email = helpers.Santize(email)

		var find bson.M
		errFindUsername := authenticationCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&find)
		errFindEmail := authenticationCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&find)

		if errFindEmail == nil || errFindUsername == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Account already exists"})
			return
		}

		password, err := helpers.Hash(password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		newAccount := bson.M{
			"username": username,
			"password": password,
			"email":    email,
		}

		result, err := authenticationCollection.InsertOne(context.TODO(), newAccount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return

		}
		c.JSON(http.StatusCreated, gin.H{"data": result})
	}
}

func LoginAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if govalidator.IsNull(username) || govalidator.IsNull(password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data can not empty"})
			return
		}

		username = helpers.Santize(username)
		password = helpers.Santize(password)

		var find bson.M
		err := authenticationCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&find)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//Convert interface to string
		hashedPassword := fmt.Sprintf("%v", find["password"])
		err = helpers.CheckPasswordHash(hashedPassword, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		token, errCreate := helpers.CreateJWT(username)
		if errCreate != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errCreate.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": token})

	}
}

func LogoutAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := authenticationCollection.FindOne(context.TODO(), bson.M{"username": c.PostForm("username")})

		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": result})

	}
}
