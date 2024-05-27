package cart

import "bookstore/dao/cart"

type CartService struct {
	cartRepo *cart.CartRepository
}

func NewCartRepository() *CartService {
	return &CartService{
		cartRepo: cart.NewCartRepository(),
	}
}
