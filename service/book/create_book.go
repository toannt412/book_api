package book

import (
	"bookstore/serialize"
	"context"
)

func (s *BookService) CreateBook(ctx context.Context, newBook *serialize.Book) (*serialize.Book, error) {
	result, err := s.bookRepo.CreateBook(ctx, newBook)
	if err != nil {
		return nil, err
	}
	categories := make([]serialize.Category, len(result.CategoryIDs))
	for i, cate := range result.CategoryIDs {
		categories[i] = serialize.Category{
			Id:      cate.Id,
			CatName: cate.CatName,
		}
	}

	authors := make([]serialize.Author, len(result.AuthorID))
	for i, author := range result.AuthorID {
		authors[i] = serialize.Author{
			Id:          author.Id,
			AuthorName:  author.AuthorName,
			DateOfBirth: author.DateOfBirth,
			HomeTown:    author.HomeTown,
			Alive:       author.Alive,
		}
	}
	return &serialize.Book{
		Id:                result.Id,
		BookName:          result.BookName,
		Price:             result.Price,
		PublishingCompany: result.PublishingCompany,
		PublicationDate:   result.PublicationDate,
		Description:       result.Description,
		CategoryIDs:       categories,
		AuthorID:          authors,
	}, nil
}
