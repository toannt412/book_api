package category

import (
	"bookstore/serialize"
	"context"
)

func (s *CategoryService) GetCategoryByID(ctx context.Context, categoryID string) (*serialize.Category, error) {
	result, err := s.categoryRepo.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	return &serialize.Category{
		Id:      result.Id,
		CatName: result.CatName,
	}, nil
}

func (s *CategoryService) GetAllCategories(ctx context.Context) ([]*serialize.Category, error) {
	result, err := s.categoryRepo.GetAllCategories(ctx)
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
