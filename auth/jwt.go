package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("could not parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token is expired")
		return
	}
	return
}

func GetUsernameFromToken(signedToken string) (username string, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("could not parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token is expired")
		return
	}
	return claims.Username, nil
}

func CheckValidToken(signedToken string) (err error) {
	_, err = jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return err
	}
	return nil
}
func RemoveToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil // Thay "your-secret-key" bằng khóa bí mật thực tế của bạn
		})
	// token, err := jwt.ParseWithClaims(
	// 	signedToken,
	// 	&JWTClaim{},
	// 	func(token *jwt.Token) (interface{}, error) {
	// 		return []byte(jwtKey), nil
	// 	},
	// )
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		claims.ExpiresAt = time.Now().Add(-1 * time.Hour).Unix()
		return
	}
	return errors.New("token is expired")
}
