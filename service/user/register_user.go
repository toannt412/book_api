package user

import (
	"context"
)

func (s *UserService) RegisterAccount(ctx context.Context, username, password, email, phone string) (string, error) {
	result, err := s.userRepo.RegisterAccount(ctx, username, password, email, phone)
	if err != nil {
		return "Register Fail", err
	}
	return result, nil
}
