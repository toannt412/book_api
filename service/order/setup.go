package order

import "bookstore/dao/cart"

type OrderService struct {
	orderService *cart.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderService: cart.NewOrderRepository(),
	}
}
