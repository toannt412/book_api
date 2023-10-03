package book

import (
	"bookstore/dao/book"
	"bookstore/dao/book/model"
	"context"
)

func CreateBook(ctx context.Context, newBook model.Book) (model.Book, error) {
	result, err := book.CreateBook(ctx, newBook)
	if err != nil {
		return model.Book{}, err
	}
	return result, nil
}
