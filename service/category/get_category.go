package category

import (
	"bookstore/dao/book"
	"bookstore/serialize"
	"context"
)

func GetCategoryByID(ctx context.Context, categoryID string) (*serialize.Category, error) {
	result, err := book.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	return &serialize.Category{
		Id:      result.Id,
		CatName: result.CatName,
	}, nil
}

func GetAllCategories(ctx context.Context) ([]*serialize.Category, error) {
	result, err := book.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	categories := make([]*serialize.Category, len(result))
	for i, category := range result {
		categories[i] = &serialize.Category{
			Id:      category.Id,
			CatName: category.CatName,
		}
	}
	return categories, nil
}
