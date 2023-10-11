package service

import (
	"bookstore/helpers"
	"context"

	"github.com/pkg/errors"
)

func (s *AdminService) Login(ctx context.Context, username, password string) (string, error) {
	user, token, err := s.adminRepo.LoginAccountAdmin(ctx, username, password)
	if err != nil {
		return "Login Fail", err
	}

	user, err = s.adminRepo.GetAdminByUserName(ctx, username)

	checkPass := helpers.CheckPasswordHash(user.Password, password)
	if checkPass == nil {
		return token, nil
	}

	return "", errors.WithMessage(err, "Login Fail")
}
