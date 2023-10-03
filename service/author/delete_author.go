package author

import (
	"bookstore/dao/book"
	"context"
)

func DeleteAuthor(ctx context.Context, authorID string) (string, error) {
	result, err := book.DeleteAuthor(ctx, authorID)
	if err != nil {
		return "", err
	}
	return result, nil
}