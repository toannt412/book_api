package cart

import (
	"bookstore/serialize"
	"context"
)

func (s *CartService) GetCart(cxt context.Context, cartID string) (*serialize.Cart, error) {
	result, err := s.cartRepo.GetCart(cxt, cartID)
	if err != nil {
		return nil, err
	}

	bookSlice := make([]serialize.CartBook, len(result.Books))
	for i, book := range result.Books {
		bookSlice[i] = serialize.CartBook{
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
		Books:         bookSlice,
		TotalQuantity: result.TotalQuantity,
		TotalAmount:   result.TotalAmount,
	}, nil
}
