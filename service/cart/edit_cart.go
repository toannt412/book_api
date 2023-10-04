package cart

import (
	"bookstore/dao/cart"
	"bookstore/serialize"
	"context"
)

func EditCart(cxt context.Context, cartID string, updateCart *serialize.Cart) (*serialize.Cart, error) {
	result, err := cart.EditCart(cxt, cartID, updateCart)
	if err != nil {
		return nil, err
	}

	cartslice := make([]serialize.CartBook, len(result.Books))
	for i, book := range result.Books {
		cartslice[i] = serialize.CartBook{
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
		Books:         cartslice,
		TotalQuantity: result.TotalQuantity,
		TotalAmount:   result.TotalAmount,
	}, nil
}
