package service

import (
	"bookstore/dao/admin"
	"bookstore/dao/admin/model"
	"bookstore/serialize"
	"context"
)

func GetAdminByID(ctx context.Context, id string) (*serialize.Admin, error) {
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

func GetAdminByEmail(ctx context.Context, email string) error {
	_, err := admin.GetAdminByEmail(ctx, email)
	if err != nil {
		return err
	}
	return nil
}

func GetAdminByUserName(ctx context.Context, username string) error {
	_, err := admin.GetAdminByUserName(ctx, username)
	if err != nil {
		return err
	}
	return nil
}

func GetAllAdmins(ctx context.Context) ([]model.Admin, error) {
	admins, err := admin.GetAllAdmins(ctx)
	if err != nil {
		return nil, err
	}
	return admins, nil
}
