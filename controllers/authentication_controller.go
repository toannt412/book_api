package controllers

import (
	"bookstore/configs"
	"bookstore/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var authenticationCollection *mongo.Collection = configs.GetCollection(configs.DB, "authentications")

// 		c.JSON(http.StatusCreated, gin.H{"data": result})
// 	}
// }

func LogoutAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := authenticationCollection.FindOne(context.TODO(), bson.M{"username": c.PostForm("username")})

		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": result})

	}
}

// Update Infomation User
func EditAUserAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		accountID := c.Param("accountId")
		var account models.Account
		defer cancel()
		objt, _ := primitive.ObjectIDFromHex(accountID)

		if err := c.BindJSON(&account); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		update := models.Account{
			Id:          objt,
			UserName:    account.UserName,
			Password:    account.Password,
			Email:       account.Email,
			FullName:    account.FullName,
			Address:     account.Address,
			DateOfBirth: account.DateOfBirth,
			Phone:       account.Phone,
		}
		result, err := authenticationCollection.UpdateOne(ctx, bson.M{"_id": objt}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var updatedAccount models.Account
		if result.MatchedCount == 1 {
			err := authenticationCollection.FindOne(ctx, bson.M{"_id": objt}).Decode(&updatedAccount)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

		}
		c.JSON(http.StatusOK, gin.H{"data": updatedAccount})
	}
}

// GetUserAuth
func GetInfoAUserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		accountID := c.Param("accountId")
		var account models.Account
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(accountID)

		err := authenticationCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": account})
	}
}
