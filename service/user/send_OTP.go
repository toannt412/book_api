package user

import (
	"bookstore/helpers"
	"context"
	"errors"
)

func (s *UserService) SendOTPByPhone(ctx context.Context, phone string) (string, error) {
	otp, err := helpers.GenerateOTP()
	if err != nil {
		return "", err
	}

	if !s.userRepo.SaveOTPByPhone(ctx, phone, otp) {
		return "", errors.New("failed to save otp")
	}
	result, err := s.twilio.SendOTPByPhone(otp)
	if err != nil {
		return "", err

	}
	return result, nil
}

func (s *UserService) SendOTPByEmail(ctx context.Context, email string) (string, error) {
	otp, err := helpers.GenerateOTP()
	if err != nil {
		return "", err
	}
	if !s.userRepo.SaveOTPByEmail(ctx, email, otp) {
		return "", errors.New("failed to save otp")
	}
	result, err := s.brevo.SendOTPByEmail(otp, email)
	if err != nil {
		return "", err
	}
	return result, nil
}
