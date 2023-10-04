package cart

import (
	"bookstore/dao/cart"
	"bookstore/serialize"
	"context"
)

func CreateCart(cxt context.Context, newCart *serialize.Cart) (*serialize.Cart, error){
	result, err := cart.CreateCart(cxt, newCart)
	if err != nil {
		return nil, err
	}

	books := make([]serialize.CartBook, len(result.Books))
	for i, book := range result.Books {
		books[i] = serialize.CartBook{
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
		Books:         books,
		TotalQuantity: result.TotalQuantity,
		TotalAmount:   result.TotalAmount,
	}, nil
}

