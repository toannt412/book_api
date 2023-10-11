package user

import (
	"bookstore/helpers"
	"context"

	"github.com/pkg/errors"
)

func (s *UserService) Login(ctx context.Context, username, password string) (string, error) {
	result, token, err := s.userRepo.LoginAccount(ctx, username, password)
	if err != nil {
		return "Login Fail", err
	}

	result, err = s.userRepo.GetUserByUserName(ctx, username)

	checkPass := helpers.CheckPasswordHash(result.Password, password)
	if checkPass == nil {
		return token, nil
	}

	return "", errors.WithMessage(err, "Login Fail")
}
