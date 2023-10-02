package user

import (
	"bookstore/dao/user"
	"bookstore/helpers"
	"context"

	"github.com/pkg/errors"
)

func Login(ctx context.Context, username, password string) (string, error) {
	result, token, err := user.LoginAccount(ctx, username, password)
	if err != nil {
		return "Login Fail", err
	}

	result, err = user.GetUserByUserName(ctx, username)

	checkPass := helpers.CheckPasswordHash(result.Password, password)
	if checkPass == nil {
		return token, nil
	}

	return "", errors.WithMessage(err, "Login Fail")
}
