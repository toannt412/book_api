package service

import (
	"bookstore/dao/admin"
	"bookstore/helpers"
	"context"

	"github.com/pkg/errors"
)

func Login(ctx context.Context, username, password string) (string, error) {
	user, token, err := admin.LoginAccountAdmin(ctx, username, password)
	if err != nil {
		return "Login Fail", err
	}

	user, err = admin.GetAdminByUserName(ctx, username)

	checkPass := helpers.CheckPasswordHash(user.Password, password)
	if checkPass == nil {
		return token, nil
	}

	return "", errors.WithMessage(err, "Login Fail")
}
