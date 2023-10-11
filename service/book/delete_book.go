package book

import (
	"context"
)

func (s *BookService) DeleteBook(ctx context.Context, bookID string) (string, error) {
	result, err := s.bookRepo.DeleteBook(ctx, bookID)
	if err != nil {
		return "Deleted failed", err
	}
	return result, nil
}
