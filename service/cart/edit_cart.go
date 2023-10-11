package cart

import (
	"bookstore/serialize"
	"context"
)

func (s *CartService) EditCart(cxt context.Context, cartID string, updateCart *serialize.Cart) (*serialize.Cart, error) {
	result, err := s.cartRepo.EditCart(cxt, cartID, updateCart)
	if err != nil {
		return nil, err
	}

	cartSlice := make([]serialize.CartBook, len(result.Books))
	for i, book := range result.Books {
		cartSlice[i] = serialize.CartBook{
			BookID:   book.BookID,
			BookName: book.BookName,
			Price:    book.Price,
			Quantity: book.Quantity,
			Total:    book.Total,
		}
	}
	return &serialize.Cart{
		Id:            result.Id,
		UserID:        result.UserID,
		Books:         cartSlice,
		TotalQuantity: result.TotalQuantity,
		TotalAmount:   result.TotalAmount,
	}, nil
}
