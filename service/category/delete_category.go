package category

import (
	"bookstore/dao/book"
	"context"
)

func DeleteCategory(ctx context.Context, categoryID string) (string, error) {
	result, err := book.DeleteCategory(ctx, categoryID)
	if err != nil {
		return "Delete Failed", err
	}
	return result, nil
}