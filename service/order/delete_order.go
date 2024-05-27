package order

import (
	"context"
)

func (s *OrderService) DeleteOrder(cxt context.Context, orderID string) (string, error) {
	result, err := s.orderService.DeleteOrder(cxt, orderID)
	if err != nil {
		return "Deleted failed", err
	}
	return result, nil
}
