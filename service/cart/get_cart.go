package cart

import (
	"bookstore/dao/cart"
	"bookstore/serialize"
	"context"
)

func GetCart(cxt context.Context, cartID string) (*serialize.Cart, error) {
	result, err := cart.GetCart(cxt, cartID)
	if err != nil {
		return nil, err
	}

	bookslice := make([]serialize.CartBook, len(result.Books))
	for i, book := range result.Books {
		bookslice[i] = serialize.CartBook{
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
		Books:         bookslice,
		TotalQuantity: result.TotalQuantity,
		TotalAmount:   result.TotalAmount,
	}, nil
}
