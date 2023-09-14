package routes

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CheckToken() gin.HandlerFunc {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	requiredToken := os.Getenv("SECRET_KEY")

	if requiredToken == "" {
		log.Fatal("Error loading .env file")

	}

	//verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(requiredToken))
	return func(c *gin.Context) {

		//Read the token out of the response body
		buf := new(bytes.Buffer)
		io.Copy(buf, c.Request.Body)
		c.Request.Body.Close()
		//tokenString := strings.TrimSpace(buf.String())

		// Parse the token
		// _, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 	// since we only use the one private key to sign the tokens,
		// 	// we also only use its public counter part to verify
		// 	return verifyKey, nil
		// })
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Token invalido"})
			return
		}

		// token := c.Request.Header.Get("Authorization")
		// if token != requiredToken {
		// 	c.JSON(http.StatusBadRequest, gin.H{"message": "Token invalido"})
		// 	return
		// }
		// if token == "" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"message": "Token deve ser preenchido"})
		// 	return
		// }

		c.Next()

	}
}
