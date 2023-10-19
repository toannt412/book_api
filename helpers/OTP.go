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

func SetOTP(otp string, user *model.User) {
	//otp, _ := GenerateOTP()
	expiry := time.Now().Add(5 * time.Minute)

	user.OTP = otp
	user.OTPExpiry = expiry
}

func VerifyOTP(user *model.User, otp string) bool {
	return user.OTP == otp && time.Now().Before(user.OTPExpiry)
}
