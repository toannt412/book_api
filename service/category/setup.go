package category

import "bookstore/dao/book"

type CategoryService struct {
	categoryRepo *book.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		categoryRepo: book.NewCategoryRepository(),
	}
}
