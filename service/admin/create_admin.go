package service

import (
	"bookstore/dao/admin"
	"bookstore/dao/admin/model"
	"bookstore/serialize"
	"context"
)

func CreateAdmin(ctx context.Context, sth model.Admin) (*serialize.Admin, error) {
	user, err := admin.CreateAdmin(ctx, sth)
	if err != nil {
		return nil, err
	}
	return &serialize.Admin{
		Id:       user.Id.Hex(),
		UserName: user.UserName,
		Password: user.Password,
		FullName: user.FullName,
		Phone:    user.Phone,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}
