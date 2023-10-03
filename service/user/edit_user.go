package user

import (
	"bookstore/dao/user"
	"bookstore/serialize"
	"context"
)

func EditUser(ctx context.Context, id string, sth *serialize.User) (*serialize.User, error) {
	user, err := user.EditUser(ctx, id, sth)
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
