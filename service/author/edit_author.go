package author

import (
	"bookstore/dao/book"
	"bookstore/serialize"
	"context"
)

func EditAuthor(ctx context.Context, id string, author *serialize.Author) (*serialize.Author, error) {
	result, err := book.EditAuthor(ctx, id, author)
	if err != nil {
		return nil, err
	}
	return &serialize.Author{
		Id:          result.Id,
		AuthorName:  result.AuthorName,
		DateOfBirth: result.DateOfBirth,
		HomeTown:    result.HomeTown,
		Alive:       result.Alive,
	}, nil
}