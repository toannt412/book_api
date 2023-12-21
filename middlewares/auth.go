package middlewares

import (
	"bookstore/auth"
	service "bookstore/service/admin"
	"bookstore/service/user"

	"github.com/gin-gonic/gin"
)

type Middlewares struct {
	adminSvc *service.AdminService
	userSvc  *user.UserService
}

func NewMiddlewares() *Middlewares {
	return &Middlewares{
		adminSvc: service.NewAdminService(),
		userSvc:  user.NewUserService(),
	}
}
func (m *Middlewares) AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		_, errToken := m.adminSvc.GetAdminToken(c, tokenString)
		if errToken != nil {
			c.JSON(401, gin.H{"error": "token is invalid"})
			c.Abort()
			return

		}

		isValidToken := auth.CheckValidToken(tokenString)
		if isValidToken != nil {
			c.JSON(401, gin.H{"error": isValidToken})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (m *Middlewares) LogoutAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		isValidToken := auth.CheckValidToken(tokenString)
		if isValidToken != nil {
			c.JSON(401, gin.H{"error": isValidToken})
			c.Abort()
			return
		}
		_, checkToken := m.adminSvc.GetAdminToken(c, tokenString)
		if checkToken == nil {
			errRemoveToken := m.adminSvc.EditAdminToken(c, tokenString)
			if errRemoveToken != nil {
				c.JSON(401, gin.H{"error": errRemoveToken})
				c.Abort()
				return
			}
		}
		c.JSON(200, gin.H{"status": "logout success"})
	}
}

func (m *Middlewares) AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		_, errToken := m.userSvc.GetUserToken(c, tokenString)
		if errToken != nil {
			c.JSON(401, gin.H{"error": "token is invalid"})
			c.Abort()
			return

		}
		isValidToken := auth.CheckValidToken(tokenString)
		if isValidToken != nil {
			c.JSON(401, gin.H{"error": isValidToken.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (m *Middlewares) LogoutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		isValidToken := auth.CheckValidToken(tokenString)
		if isValidToken != nil {
			c.JSON(401, gin.H{"error": isValidToken})
			c.Abort()
			return
		}

		_, checkToken := m.userSvc.GetUserToken(c, tokenString)
		if checkToken == nil {
			errRemoveToken := m.userSvc.EditUserToken(c, tokenString)
			if errRemoveToken != nil {
				c.JSON(401, gin.H{"error": errRemoveToken})
				c.Abort()
				return
			}
		}
		c.JSON(200, gin.H{"status": "logout success"})
	}
}
