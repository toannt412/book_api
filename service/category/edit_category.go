package category

import (
	"bookstore/serialize"
	"context"
)

func (s *CategoryService) EditCategory(ctx context.Context, categoryID string, category *serialize.Category) (*serialize.Category, error) {
	result, err := s.categoryRepo.EditCategory(ctx, categoryID, category)
	if err != nil {
		return nil, err
	}
	return &serialize.Category{
		Id:      result.Id,
		CatName: result.CatName,
	}, nil
}
