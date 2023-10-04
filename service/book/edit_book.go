package book

import (
	"bookstore/dao/book"
	"bookstore/serialize"
	"context"
)

func EditBook(ctx context.Context, bookID string, updateBook *serialize.Book) (*serialize.Book, error) {
	result, err := book.EditBook(ctx, bookID, updateBook)
	if err != nil {
		return &serialize.Book{}, err
	}

	categoryIDs := make([]serialize.Category, len(result.CategoryIDs))
	for i, category := range result.CategoryIDs {
		categoryIDs[i] = serialize.Category(category)
	}
	authorSlice := make([]serialize.Author, len(result.AuthorID))
	for i, author := range result.AuthorID {
		authorSlice[i] = serialize.Author(author)
	}
	return &serialize.Book{
		Id:                result.Id,
		BookName:          result.BookName,
		Price:             result.Price,
		PublishingCompany: result.PublishingCompany,
		PublicationDate:   result.PublicationDate,
		Description:       result.Description,
		CategoryIDs:       categoryIDs,
		AuthorID:          authorSlice,
	}, nil
}
