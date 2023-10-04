package order

import (
	"bookstore/dao/cart"
	"bookstore/serialize"
	"context"
)

func GetOrderByID(cxt context.Context, orderID string) (*serialize.Order, error) {
	result, err := cart.GetOrderByID(cxt, orderID)
	if err != nil {
		return nil, err
	}

	bookSlice := make([]serialize.OrderBook, len(result.Books))
	for i, book := range result.Books {
		bookSlice[i] = serialize.OrderBook{
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
		Books:         bookSlice,
		TotalQuantity: result.TotalQuantity,
		TotalPrice:    result.TotalPrice,
		TotalAmount:   result.TotalAmount,
		Status:        result.Status,
		// CartID:        result.CartID,
	}, nil
}
