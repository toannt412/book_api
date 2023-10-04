package category

import (
	"bookstore/dao/book"
	"bookstore/serialize"
	"context"
)

func CreateCategory(ctx context.Context, newCategory *serialize.Category) (*serialize.Category, error) {
	ressult, err := book.CreateCategory(ctx, newCategory)
	if err != nil {
		return nil, err
	}
	return &serialize.Category{
		Id:      ressult.Id,
		CatName: ressult.CatName,
	}, nil
}