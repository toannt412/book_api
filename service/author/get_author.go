package author

import (
	"bookstore/serialize"
	"context"
)

func (s *AuthorService) GetAuthorByID(ctx context.Context, authorID string) (*serialize.Author, error) {
	result, err := s.authorRepo.GetAuthorByID(ctx, authorID)
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

func (s *AuthorService) GetAllAuthors(ctx context.Context) ([]serialize.Author, error) {
	result, err := s.authorRepo.GetAllAuthors(ctx)
	if err != nil {
		return nil, err
	}

	allAuthors := make([]serialize.Author, len(result))
	for i, author := range result {
		allAuthors[i] = serialize.Author{
			Id:          author.Id,
			AuthorName:  author.AuthorName,
			DateOfBirth: author.DateOfBirth,
			HomeTown:    author.HomeTown,
			Alive:       author.Alive,
		}
	}
	return allAuthors, nil
}
