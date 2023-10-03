package user

import (
	"bookstore/dao/user"
	"context"
)

func DeleteUser(ctx context.Context, id string) (string, error) {
	res, err := user.DeleteUser(ctx, id)
	if err != nil {
		return "Deleted fail", err
	}
	return res, nil
}