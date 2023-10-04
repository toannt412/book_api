package order

import (
	"bookstore/dao/cart"
	"bookstore/serialize"
	"context"
)

func CreateOrder(cxt context.Context, newOrder *serialize.Order) (*serialize.Order, error) {
	result, err := cart.CreateOrder(cxt, newOrder)
	if err != nil {
		return nil, err
	}
	books := make([]serialize.OrderBook, len(result.Books))
	for i, book := range result.Books {
		books[i] = serialize.OrderBook{
			BookID:   book.BookID,
			BookName: book.BookName,
			Price:    book.Price,
			Quantity: book.Quantity,
			Total:    book.Total,
		}
	}
	return &serialize.Order{
		Id:            result.Id,
		UserID:        result.UserID,
		Books:         books,
		CartID:        result.CartID,
		TotalQuantity: result.TotalQuantity,
		TotalPrice:    result.TotalPrice,
		TotalAmount:   result.TotalAmount,
		Status:        result.Status,
	}, nil
}
