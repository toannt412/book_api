package user

import (
	"bookstore/dao/user"
	"context"
)

func RegisterAccount(ctx context.Context, username, password, email string) (string, error) {
	result, err := user.RegisterAccount(ctx, username, password, email)
	if err != nil {
		return "Register Fail", err
	}
	return result, nil
}