package book

import (
	"bookstore/dao/book"
	"context"
)

func DeleteBook(ctx context.Context, bookID string) (string, error) {
	result, err := book.DeleteBook(ctx, bookID)
	if err != nil {
		return "Deleted failed", err
	}
	return result, nil
}
