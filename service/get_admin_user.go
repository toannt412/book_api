package service

import (
	"bookstore/dao/admin"
	"bookstore/serialize"
	"context"

	"github.com/pkg/errors"
)

func GetAdminUserByID(ctx context.Context, id string) (*serialize.Admin, error) {
	user, err := admin.GetAdminByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &serialize.Admin{
		Id:       user.Id.Hex(),
		UserName: user.UserName,
		FullName: user.FullName,
		Phone:    user.Phone,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func GetAdminByUserName(ctx context.Context, username string) (*serialize.Admin, error) {
	user, err := admin.GetAdminByUserName(ctx, username)
	if err != nil {
		return nil, err
	}
	return &serialize.Admin{
		UserName: user.UserName,
		FullName: user.FullName,
	}, nil
}

func Login(ctx context.Context, id, password string) (string, error) {
	token := "anbfdsfi"

	user, err := admin.GetAdminByID(ctx, id)

	if user.Password == password {
		return token, nil
	}

	return "", errors.WithMessage(err, "Sai mat khau roi")
}
