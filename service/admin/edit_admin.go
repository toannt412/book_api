package service

import (
	"bookstore/serialize"
	"bookstore/dao/admin"
	"context"
)

func EditAdmin(ctx context.Context, id string, sth *serialize.Admin) (*serialize.Admin, error) {
	user, err := admin.EditAdmin(ctx, id, sth)
	if err != nil {
		return nil, err
	}
	return &serialize.Admin{
		FullName: user.FullName,
		Phone:    user.Phone,
		Role:     user.Role,
	}, nil
}
