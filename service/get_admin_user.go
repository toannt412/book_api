package service

import (
	"bookstore/dao/admin"
	"bookstore/serialize"
	"context"
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


