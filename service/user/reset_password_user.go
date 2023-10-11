package user

import (
	"context"
)

func (s *UserService) ResetPassword(ctx context.Context, password string) (string, error) {
	result, err := s.userRepo.ResetPassword(ctx, password)
	if err != nil {
		return "", err
	}
	return result, nil
}
