package category

import (
	"bookstore/dao/book"
	"bookstore/serialize"
	"context"
)

func EditCategory(ctx context.Context, categoryID string, category *serialize.Category) (*serialize.Category, error) {
	result, err := book.EditCategory(ctx, categoryID, category)
	if err != nil {
		return nil, err
	}
	return &serialize.Category{
		Id:      result.Id,
		CatName: result.CatName,
	}, nil
}
