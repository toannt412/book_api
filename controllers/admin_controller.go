package controllers

import (
	"bookstore/dao/admin/model"
	"bookstore/helpers"
	"bookstore/responses"
	"bookstore/serialize"
	service "bookstore/service/admin"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//var adminsCollection *mongo.Collection = configs.GetCollection(configs.DB, "admins")

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

		res, err := service.Login(c, username, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "Username or password is incorrect"})
			return
		}
		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "Login Success", Data: map[string]interface{}{"token": res}})

	}
}

func GetAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminId := c.Param("adminId")

		res, err := service.GetAdminByID(c, adminId)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})

	}
}

func EditAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminId := c.Param("adminId")
		var admin *serialize.Admin

		if err := c.BindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		update := &serialize.Admin{
			FullName: admin.FullName,
			Phone:    admin.Phone,
			Role:     admin.Role,
		}
		res, err := service.EditAdmin(c, adminId, update)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

func DeleteAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminId := c.Param("adminId")

		res, err := service.DeleteAdmin(c, adminId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"message": res}})
	}
}

func GetAllAdmins() gin.HandlerFunc {
	return func(c *gin.Context) {

		admins, err := service.GetAllAdmins(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admins}})
	}
}

func CreateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Validate the request body
		var admin model.Admin
		if err := c.BindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newAdmin := model.Admin{
			Id:       primitive.NewObjectID(),
			FullName: admin.FullName,
			Phone:    admin.Phone,
			Role:     admin.Role,
			UserName: admin.UserName,
			Password: admin.Password,
			Email:    admin.Email,
		}

		// 		// 		// Generate a new MongoDB ObjectID
		// 		username := c.PostForm("username")
		// 		password := c.PostForm("password")
		// 		email := c.PostForm("email")
		// 		role := c.PostForm("role")

		// 		// 		// if username == "" || password == "" || email == "" || role == "" {
		// 		// 		// 	c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Missing username, password, role or email"}})
		// 		// 		// 	return
		// 		// 		// } //gin.H{"error": "Missing username, password, role or email"

		if govalidator.IsNull(admin.UserName) || govalidator.IsNull(admin.Email) || govalidator.IsNull(admin.Password) || govalidator.IsNull(admin.Role) {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Missing username, password, role or email"}})
			return
		}

		if !govalidator.IsEmail(admin.Email) {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Email is invalid"}})
			return
		}

		admin.UserName = helpers.Santize(admin.UserName)
		admin.Password = helpers.Santize(admin.Password)
		admin.Email = helpers.Santize(admin.Email)

		errFindEmail := service.GetAdminByEmail(c, admin.Email)
		errFindUsername := service.GetAdminByUserName(c, admin.UserName)
		if errFindEmail == nil || errFindUsername == nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Account already exists"}})
			return
		}

		// jsData := &serialize.Admin{
		// 	Id:       newAdmin.Id.Hex(),
		// 	UserName: newAdmin.UserName,
		// 	Password: newAdmin.Password,
		// 	FullName: newAdmin.FullName,
		// 	Phone:    newAdmin.Phone,
		// 	Email:    newAdmin.Email,
		// 	Role:     newAdmin.Role,
		// }

		res, err := service.CreateAdmin(c, newAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.AdminResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": res}})
		// 		var find bson.M
		// 		errFindUsername := adminsCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&find)
		// 		errFindEmail := adminsCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&find)

		// 		if errFindEmail == nil || errFindUsername == nil {
		// 			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Account already exists"}})
		// 			return
		// 		}

		// 		password, err := helpers.Hash(password)
		// 		if err != nil {
		// 			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		// 			return
		// 		}

		// 		// 		newAdmin := bson.M{
		// 		// 			"username": username,
		// 		// 			"password": password,
		// 		// 			"email":    email,
		// 		// 			"role":     role,
		// 		// 		}

		// 		// 		// Insert the admin document into the database
		// 		// 		result, err := adminsCollection.InsertOne(ctx, newAdmin)
		// 		// 		if err != nil {
		// 		// 			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		// 		// 			return
		// 		// 		}

		// 		// 		c.JSON(http.StatusCreated, responses.AdminResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
