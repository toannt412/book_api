package helpers

import (
	"bookstore/dao/user/model"
	"crypto/rand"
	"time"
)

const otpChars = "1234567890"

func GenerateOTP() (string, error) {
	buffer := make([]byte, 6)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < 6; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}
	return string(buffer), nil
}

func VerifyOTP(user *model.User, otp string) bool {
	return user.OTP == otp && time.Now().Before(user.OTPExpiry)
}
