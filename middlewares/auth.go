package middlewares

import (
	"bookstore/configs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var adminsCollection *mongo.Collection = configs.GetCollection(configs.DB, "admins")
var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		err := adminsCollection.FindOne(c, bson.M{"token": tokenString})
		if err.Err() != nil {
			c.JSON(401, gin.H{"error": "token is invalid"})
			c.Abort()
			return
		}
		// err := auth.ValidateToken(tokenString)
		// if err != nil {
		// 	c.JSON(401, gin.H{"error": err.Error()})
		// 	c.Abort()
		// 	return
		// }
		c.Next()
	}
}

func LogoutAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		var find bson.M
		checkToken := adminsCollection.FindOne(c, bson.M{"token": tokenString}).Decode(&find)
		if checkToken == nil {
			filter := bson.M{"_id": find["_id"]}
			update := bson.M{"$unset": bson.M{"token": ""}}
			_, err := adminsCollection.UpdateOne(c, filter, update)
			if err != nil {
				c.JSON(401, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}
		c.JSON(200, gin.H{"status": "logout success"})
	}
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		err := userCollection.FindOne(c, bson.M{"token": tokenString})
		if err.Err() != nil {
			c.JSON(401, gin.H{"error": "token is invalid"})
			c.Abort()
			return
		}
		// err := auth.ValidateToken(tokenString)
		// if err != nil {
		// 	c.JSON(401, gin.H{"error": err.Error()})
		// 	c.Abort()
		// 	return
		// }
		c.Next()
	}
}

func LogoutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		var find bson.M
		checkToken := userCollection.FindOne(c, bson.M{"token": tokenString}).Decode(&find)
		if checkToken == nil {
			filter := bson.M{"_id": find["_id"]}
			update := bson.M{"$unset": bson.M{"token": ""}}
			_, err := userCollection.UpdateOne(c, filter, update)
			if err != nil {
				c.JSON(401, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}
		c.JSON(200, gin.H{"status": "logout success"})
	}
}
