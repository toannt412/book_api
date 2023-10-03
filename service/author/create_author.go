package author

import (
	"bookstore/dao/book"
	"bookstore/serialize"
	"context"
)

func CreateAuthor(ctx context.Context, newAuthor *serialize.Author) (*serialize.Author, error) {
	result, err := book.CreateAuthor(ctx, newAuthor)
	if err != nil{
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