package user

import (
	"context"
)

func (s *UserService) ForgotPassword(ctx context.Context, phone string) (string, error) {
	result, err := s.userRepo.ForgotPassword(ctx, phone)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *UserService) ResetPassword(ctx context.Context, otp, password string) (string, error) {
	result, err := s.userRepo.ResetPassword(ctx, otp, password)
	if err != nil {
		return "", err
	}
	return result, nil
}
