package order

import (
	"bookstore/dao/cart"
	"context"
)

func DeleteOrder(cxt context.Context, orderID string) (string, error) {
	result, err := cart.DeleteOrder(cxt, orderID)
	if err != nil {
		return "Deleted failed", err
	}
	return result, nil
}
