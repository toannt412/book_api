package book

import (
	"bookstore/dao/book"
	"bookstore/dao/book/model"
	"context"
)

func GetAllBooks(ctx context.Context) ([]model.Book, error) {
	result, err := book.GetAllBooks(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetBookByID(ctx context.Context, bookID string) (model.Book, error) {
	result, err := book.GetBookByID(ctx, bookID)
	if err != nil {
		return model.Book{}, err
	}
	return result, nil
}