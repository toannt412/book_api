package user

import (
	"context"
)

func (s *UserService) Login(ctx context.Context, username, password string) (string, error) {
	token, err := s.userRepo.LoginAccount(ctx, username, password)
	if err != nil {
		return "wrong username or password", err
	}
	return token, nil

}
