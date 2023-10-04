package cart

import (
	"bookstore/dao/cart"
	"context"
)
func DeleteCart(cxt context.Context, cartID string) (string, error) {
	result, err := cart.DeleteCart(cxt, cartID)
	if err != nil {
		return "Deleted failed", err
	}
	return result, nil
}