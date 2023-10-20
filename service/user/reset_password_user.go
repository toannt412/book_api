package user

import (
	"bookstore/helpers"
	"context"
	"errors"
)

func (s *UserService) ResetPassword(ctx context.Context, otp, password, phone string) (string, error) {
	user, _ := s.userRepo.GetUserByPhone(ctx, phone)
	if !helpers.VerifyOTP(&user, otp) {
		return "", errors.New("invalid otp")
	}
	if !s.userRepo.ResetPassword(ctx, otp, password){
		return "", errors.New("failed to reset password")
	}
	
	return "reset password success", nil
}
