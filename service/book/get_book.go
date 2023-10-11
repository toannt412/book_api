package book

import (
	"bookstore/dao/book/model"
	"context"
)

func (s *BookService) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	result, err := s.bookRepo.GetAllBooks(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *BookService) GetBookByID(ctx context.Context, bookID string) (model.Book, error) {
	result, err := s.bookRepo.GetBookByID(ctx, bookID)
	if err != nil {
		return model.Book{}, err
	}
	return result, nil
}
