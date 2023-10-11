package author

import "bookstore/dao/book"

type AuthorService struct {
	authorRepo *book.AuthorRepository
}

func NewAuthorService() *AuthorService {
	return &AuthorService{
		authorRepo: book.NewAuthorRepository(),
	}
}