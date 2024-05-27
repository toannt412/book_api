package category

import (
	"context"
)

func (s *CategoryService) DeleteCategory(ctx context.Context, categoryID string) (string, error) {
	result, err := s.categoryRepo.DeleteCategory(ctx, categoryID)
	if err != nil {
		return "Delete Failed", err
	}
	return result, nil
}
