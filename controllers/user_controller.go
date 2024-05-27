package controllers

import (
	"bookstore/helpers"
	"bookstore/responses"
	"bookstore/serialize"
	service "bookstore/service/user"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSvc *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userSvc: service.NewUserService(),
	}
}

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

// READ
// GET BY ID
func (ctrl *UserController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		res, err := ctrl.userSvc.GetUserByID(c, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

// UPDATE
func (ctrl *UserController) EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		var user *serialize.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		update := &serialize.User{
			FullName:    user.FullName,
			Location:    user.Location,
			DateOfBirth: user.DateOfBirth,
			Phone:       user.Phone,
		}
		res, err := ctrl.userSvc.EditUser(c, userId, update)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

// DELETE
func (ctrl *UserController) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		res, err := ctrl.userSvc.DeleteUser(c, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
		)
	}
}

// GET ALL
func (ctrl *UserController) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := ctrl.userSvc.GetAllUsers(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
		)
	}
}

// // Register
func (ctrl *UserController) RegisterAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")
		phone := c.PostForm("phone")

		if govalidator.IsNull(username) || govalidator.IsNull(email) || govalidator.IsNull(password) || govalidator.IsNull(phone) {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "All fields are required"}})
			return
		}

		if !govalidator.IsEmail(email) {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid email"}})
			return
		}

		err := helpers.IsValidatePhoneNumber(phone)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		username = helpers.Santize(username)
		password = helpers.Santize(password)
		email = helpers.Santize(email)
		phone = helpers.Santize(phone)

		_, errFindUsername := ctrl.userSvc.GetUserByUserName(c, username)
		_, errFindEmail := ctrl.userSvc.GetUserByEmail(c, email)

		if errFindEmail == nil || errFindUsername == nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Username or email already exists"}})
			return
		}

		res, err := ctrl.userSvc.RegisterAccount(c, username, password, email, phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

// Login
func (ctrl *UserController) LoginAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if govalidator.IsNull(username) || govalidator.IsNull(password) {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "All fields are required"}})
			return
		}

		username = helpers.Santize(username)
		password = helpers.Santize(password)
		user, err := ctrl.userSvc.GetUserByUserName(c, username)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "wrong username or password"}})
			return
		}
		checkPass := helpers.CheckPasswordHash(user.Password, password)
		if checkPass != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "wrong username or password"}})
			return
		}
		token, err := ctrl.userSvc.Login(c, username, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": token}})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": token})

	}
}

func (ctrl *UserController) ResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.PostForm("password")
		otp := c.PostForm("otp")
		phone := c.PostForm("phone")
		if govalidator.IsNull(otp) || govalidator.IsNull(password) || govalidator.IsNull(phone) {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "All fields are required"}})
			return
		}

		res, err := ctrl.userSvc.ResetPassword(c, otp, password, phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})
	}
}

func (ctrl *UserController) ForgotPasswordUsePhone() gin.HandlerFunc {
	return func(c *gin.Context) {
		phone := c.PostForm("phone")

		if govalidator.IsNull(phone) {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "All fields are required"}})
			return
		}
		errPhone := helpers.IsValidatePhoneNumber(phone)
		if errPhone != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": errPhone.Error()}})
			return
		}
		_, checkPhone := ctrl.userSvc.GetUserByPhone(c, phone)
		if checkPhone != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Phone is not registered"}})
			return
		}
		res, err := ctrl.userSvc.SendOTPByPhone(c, phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"message": "Please check your SMS", "data": res}})
	}
}

func (ctrl *UserController) ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		phone := c.PostForm("phone")
		password := c.PostForm("password")

		if govalidator.IsNull(phone) || govalidator.IsNull(password) {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "All fields are required"}})
			return
		}
		errPhone := helpers.IsValidatePhoneNumber(phone)
		if errPhone != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": errPhone.Error()}})
			return
		}
		user, checkPhone := ctrl.userSvc.GetUserByPhone(c, phone)
		if checkPhone != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Phone is not registered"}})
			return
		}
		checkPass := helpers.CheckPasswordHash(user.Password, password)
		if checkPass != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "wrong password"}})
			return
		}
		res, err := ctrl.userSvc.SendOTPByPhone(c, phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})

	}
}

func (ctrl *UserController) ForgotPasswordUseEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		if govalidator.IsNull(email) {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "All fields are required"}})
			return
		}
		if !helpers.IsValidateEmail(email) {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid email"}})
			return
		}
		_, checkEmail := ctrl.userSvc.GetUserByEmail(c, email)
		if checkEmail != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Email is not registered"}})
			return
		}
		res, err := ctrl.userSvc.SendOTPByEmail(c, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})

	}
}

func (ctrl *UserController) UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, handler, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer file.Close()
		fileName := handler.Filename
		errUpload := ctrl.userSvc.UploadImage(c, file, fileName)
		if errUpload != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errUpload.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "upload success"}})

	}
}

func (ctr *UserController) GetImageFromBucket() gin.HandlerFunc {
	return func(c *gin.Context) {
		objectKey := c.Param("objectKey")
		request, errDownload := ctr.userSvc.GetPresignedURL(objectKey)
		if errDownload != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errDownload.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": request}})
	}
}

func (ctr *UserController) UploadImageUsePresignedURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		nameObject := c.Request.FormValue("objectKey")
		objectKey := nameObject + "_" + helpers.GenerateStringName()
		request, errDownload := ctr.userSvc.UploadImageUsePresignedURL(objectKey)
		if errDownload != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errDownload.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": request}})
	}
}
