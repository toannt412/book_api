package service

import (
	"bookstore/dao/admin"
	"context"
)

func DeleteAdmin(ctx context.Context, id string) (string, error) {
	res, err := admin.DeleteAdmin(ctx, id)
	if err != nil {
		return "Deleted fail", err
	}
	return res, nil
}
