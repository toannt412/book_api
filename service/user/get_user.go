package user

import (
	"bookstore/dao/user"
	"bookstore/dao/user/model"
	"bookstore/serialize"
	"context"
)

func GetUserByID(ctx context.Context, id string) (*serialize.User, error) {
	user, err := user.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &serialize.User{
		Id:          user.Id,
		FullName:    user.FullName,
		Location:    user.Location,
		DateOfBirth: user.DateOfBirth,
		Phone:       user.Phone,
	}, nil
}

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

func GetAllUsers(ctx context.Context) ([]model.User, error) {
	users, err := user.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
