package controllers

import (
	"bookstore/configs"
	"bookstore/helpers"
	"bookstore/models"
	"bookstore/responses"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var adminsCollection *mongo.Collection = configs.GetCollection(configs.DB, "admins")

func LoginAccountAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if govalidator.IsNull(username) || govalidator.IsNull(password) {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "Username or password is empty"})
			return
		}

		username = helpers.Santize(username)
		password = helpers.Santize(password)

		var find bson.M
		err := adminsCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&find)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "Username or password is incorrect"})
			return
		}

		//Convert interface to string
		hashedPassword := fmt.Sprintf("%v", find["password"])
		err = helpers.CheckPasswordHash(hashedPassword, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "Username or password is incorrect"})
			return
		}

		token, errCreate := helpers.CreateJWT(username)
		if errCreate != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "Create Token successfully", Data: map[string]interface{}{"token": token}})

	}
}

func GetAAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminId")
		var admin models.Admin
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		err := adminsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&admin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admin}})
	}
}

func EditAAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminId")
		var admin models.Admin
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		// Validate the request body
		if err := c.BindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		update := models.Admin{
			Id:       objId,
			FullName: admin.FullName,
			Phone:    admin.Phone,
			Role:     admin.Role,
		}

		_, err := adminsCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedAdmin models.Admin
		err = adminsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedAdmin}})
	}
}

func DeleteAAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		_, err := adminsCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"message": "Admin deleted successfully"}})
	}
}

func GetAllAdmins() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var admins []models.Admin

		cursor, err := adminsCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var admin models.Admin
			if err := cursor.Decode(&admin); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			admins = append(admins, admin)
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admins}})
	}
}

func CreateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Validate the request body
		//var admin models.Admin
		// if err := c.BindJSON(&admin); err != nil {
		// 	c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		// Generate a new MongoDB ObjectID
		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")
		role := c.PostForm("role")

		// if username == "" || password == "" || email == "" || role == "" {
		// 	c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Missing username, password, role or email"}})
		// 	return
		// } //gin.H{"error": "Missing username, password, role or email"

		if govalidator.IsNull(username) || govalidator.IsNull(email) || govalidator.IsNull(password) || govalidator.IsNull(role) {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Missing username, password, role or email"}})
			return
		}

		if !govalidator.IsEmail(email) {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Email is invalid"}})
			return
		}

		username = helpers.Santize(username)
		password = helpers.Santize(password)
		email = helpers.Santize(email)

		var find bson.M
		errFindUsername := adminsCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&find)
		errFindEmail := adminsCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&find)

		if errFindEmail == nil || errFindUsername == nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Account already exists"}})
			return
		}

		password, err := helpers.Hash(password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newAdmin := bson.M{
			"username": username,
			"password": password,
			"email":    email,
			"role":     role,
		}

		// Insert the admin document into the database
		result, err := adminsCollection.InsertOne(ctx, newAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AdminResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
