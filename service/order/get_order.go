package order

import (
	"bookstore/serialize"
	"context"
)

func (s *OrderService) GetOrderByID(cxt context.Context, orderID string) (*serialize.Order, error) {
	result, err := s.orderService.GetOrderByID(cxt, orderID)
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
		OrderDate:     result.OrderDate,
		CartID:        result.CartID,
	}, nil
}
