package category

import (
	"bookstore/serialize"
	"context"
)

func (s *CategoryService) CreateCategory(ctx context.Context, newCategory *serialize.Category) (*serialize.Category, error) {
	result, err := s.categoryRepo.CreateCategory(ctx, newCategory)
	if err != nil {
		return nil, err
	}
	return &serialize.Category{
		Id:      result.Id,
		CatName: result.CatName,
	}, nil
}
