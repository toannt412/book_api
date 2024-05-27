package author

import (
	"context"
)

func (s *AuthorService) DeleteAuthor(ctx context.Context, authorID string) (string, error) {
	result, err := s.authorRepo.DeleteAuthor(ctx, authorID)
	if err != nil {
		return "", err
	}
	return result, nil
}
