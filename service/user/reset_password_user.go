package user

import (
	"bookstore/configs"
	"bookstore/helpers"
	"context"
	"encoding/json"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"gopkg.in/gomail.v2"
)

func (s *UserService) ForgotPassword(ctx context.Context, phone string) (string, error) {
	otp, err := helpers.GenerateOTP()
	if err != nil {
		return "", err
	}
	errSaveOTP := s.userRepo.ForgotPassword(ctx, phone, otp)
	if err != nil {
		return "", errSaveOTP
	}
	accountSid := configs.Config.AccountSID
	authToken := configs.Config.AuthToken
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(configs.Config.ToPhone)
	params.SetFrom(configs.Config.FromPhone)
	params.SetBody("Your verification code is: " + otp)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return "", err
	} else {
		response, _ := json.Marshal(*resp)
		return string(response), nil
	}
}

func (s *UserService) ResetPassword(ctx context.Context, otp, password string) (string, error) {
	err := s.userRepo.ResetPassword(ctx, otp, password)
	if err != nil {
		return "", err
	}
	return "reset password success", nil
}

func (s *UserService) ForgotPasswordUseEmail(ctx context.Context, email string) (string, error) {
	otp, err := helpers.GenerateOTP()
	if err != nil {
		return "", err
	}
	errSaveOTP := s.userRepo.ForgotPasswordUseEmail(ctx, email, otp)
	if errSaveOTP != nil {
		return "", errSaveOTP
	}
	from := configs.Config.FromEmail
	host := configs.Config.SMTPHost
	port := 587
	apiKey := configs.Config.APIToken

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "OTP for forgot password")
	// text/html for a html email
	msg.SetBody("text/plain", "Your OTP is: "+otp)

	n := gomail.NewDialer(host, port, from, apiKey)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
	return "send email success", nil
}
