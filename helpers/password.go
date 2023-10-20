package helpers

import (
	"bookstore/configs"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	cost, err := strconv.Atoi(configs.Config.HashCost)
	if err != nil {
		return "", err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
