package author

import (
	"bookstore/serialize"
	"context"
)

func (s *AuthorService) CreateAuthor(ctx context.Context, newAuthor *serialize.Author) (*serialize.Author, error) {
	result, err := s.authorRepo.CreateAuthor(ctx, newAuthor)
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
