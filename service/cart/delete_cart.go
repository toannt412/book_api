package cart

import (
	"context"
)

func (s *CartService) DeleteCart(cxt context.Context, cartID string) (string, error) {
	result, err := s.cartRepo.DeleteCart(cxt, cartID)
	if err != nil {
		return "Deleted failed", err
	}
	return result, nil
}
