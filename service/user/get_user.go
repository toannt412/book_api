package user

import (
	"bookstore/dao/user"
	"context"
)

// func GetAdminByID(ctx context.Context, id string) (*serialize.Admin, error) {
// 	user, err := admin.GetAdminByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &serialize.Admin{
// 		Id:       user.Id.Hex(),
// 		UserName: user.UserName,
// 		FullName: user.FullName,
// 		Phone:    user.Phone,
// 		Email:    user.Email,
// 		Role:     user.Role,
// 	}, nil
// }

func GetUserByEmail(ctx context.Context, email string) error {
	_, err := user.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByUserName(ctx context.Context, username string) error {
	_, err := user.GetUserByUserName(ctx, username)
	if err != nil {
		return err
	}
	return nil
}

// func GetAllAdmins(ctx context.Context) ([]model.Admin, error) {
// 	admins, err := admin.GetAllAdmins(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return admins, nil
// }
