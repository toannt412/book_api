package book

import "bookstore/dao/book"

type BookService struct {
	bookRepo *book.BookRepository
}

func NewBookService() *BookService {
	return &BookService{
		bookRepo: book.NewBookRepository(),
	}
}