package controllers

import (
	"bookstore/helpers"
	"bookstore/responses"
	service "bookstore/service/user"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

//"github.com/go-playground/validator"

//var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

//var validate = validator.New()

// CREATE
// func CreateUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		var user models.User
// 		defer cancel()

// 		//validate the request body
// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}
// 		// TIM HIEU NAY LAI
// 		// //use the validator library to validate required fields
// 		// if validationErr := validate.Struct(&user); validationErr != nil {
// 		// 	c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
// 		// 	return
// 		// }

// 		newUser := models.User{
// 			Id:          primitive.NewObjectID(),
// 			FullName:    user.FullName,
// 			Location:    user.Location,
// 			DateOfBirth: user.DateOfBirth,
// 			Phone:       user.Phone,
// 		}

// 		result, err := userCollection.InsertOne(ctx, newUser)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
// 	}
// }

// // READ
// // GET BY ID
// func GetAUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		userId := c.Param("userId")
// 		var user models.User
// 		defer cancel()

// 		objId, _ := primitive.ObjectIDFromHex(userId)

// 		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
// 	}
// }

// UPDATE
// func EditAUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		userId := c.Param("userId")
// 		var user models.User
// 		defer cancel()
// 		objId, _ := primitive.ObjectIDFromHex(userId)

// 		//validate the request body
// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		// //use the validator library to validate required fields
// 		// if validationErr := validate.Struct(&user); validationErr != nil {
// 		//     c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
// 		//     return
// 		// }

// 		update := models.User{
// 			Id:          objId,
// 			FullName:    user.FullName,
// 			Location:    user.Location,
// 			DateOfBirth: user.DateOfBirth,
// 			Phone:       user.Phone,
// 		}
// 		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		//get updated user details
// 		var updatedUser models.User
// 		if result.MatchedCount == 1 {
// 			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 				return
// 			}

// 		}

// 		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}})
// 	}
// }

// // DELETE
// func DeleteAUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		userId := c.Param("userId")
// 		defer cancel()

// 		objId, _ := primitive.ObjectIDFromHex(userId)

// 		result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		if result.DeletedCount < 1 {
// 			c.JSON(http.StatusNotFound,
// 				responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
// 			)
// 			return
// 		}

// 		c.JSON(http.StatusOK,
// 			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
// 		)
// 	}
// }

// // GET ALL
// func GetAllUsers() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		var users []models.User
// 		defer cancel()

// 		results, err := userCollection.Find(ctx, bson.M{})

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		//reading from the db in an optimal way
// 		defer results.Close(ctx)
// 		for results.Next(ctx) {
// 			var singleUser models.User
// 			if err = results.Decode(&singleUser); err != nil {
// 				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			}

// 			users = append(users, singleUser)
// 		}

// 		c.JSON(http.StatusOK,
// 			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
// 		)
// 	}
// }

// // Register
func RegisterAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")

		if govalidator.IsNull(username) || govalidator.IsNull(email) || govalidator.IsNull(password) {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "All fields are required"}})
			return
		}

		if !govalidator.IsEmail(email) {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid email"}})
			return
		}

		username = helpers.Santize(username)
		password = helpers.Santize(password)
		email = helpers.Santize(email)

		errFindUsername := service.GetUserByUserName(c, username)
		errFindEmail := service.GetUserByEmail(c, email)

		if errFindEmail == nil || errFindUsername == nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Username or email already exists"}})
			return
		}

		res, err := service.RegisterAccount(c, username, password, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

// Login
func LoginAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if govalidator.IsNull(username) || govalidator.IsNull(password) {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "All fields are required"}})
			return
		}

		username = helpers.Santize(username)
		password = helpers.Santize(password)

		token, err := service.Login(c, username, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": token}})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": token})

	}
}
