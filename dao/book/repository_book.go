package book

import (
	"bookstore/dao"
	"bookstore/dao/book/model"
	"bookstore/serialize"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	booksCollection *mongo.Collection
}

func NewBookRepository() *BookRepository {
	var DB *mongo.Client = dao.ConnectDB()
	return &BookRepository{
		booksCollection: dao.GetCollection(DB, "books"),
	}
}

func (repo *BookRepository) CreateBook(ctx context.Context, newBook *serialize.Book) (*serialize.Book, error) {
	categorySlice := make([]model.Category, len(newBook.CategoryIDs))
	for i, cate := range newBook.CategoryIDs {
		categorySlice[i] = model.Category{
			Id:      cate.Id,
			CatName: cate.CatName,
		}
	}

	authorSlice := make([]model.Author, len(newBook.AuthorID))
	for i, author := range newBook.AuthorID {
		authorSlice[i] = model.Author{
			Id:          author.Id,
			AuthorName:  author.AuthorName,
			DateOfBirth: author.DateOfBirth,
			HomeTown:    author.HomeTown,
			Alive:       author.Alive,
		}
	}

	model := model.Book{
		Id:                newBook.Id,
		BookName:          newBook.BookName,
		Price:             newBook.Price,
		PublishingCompany: newBook.PublishingCompany,
		PublicationDate:   newBook.PublicationDate,
		Description:       newBook.Description,
		CategoryIDs:       categorySlice,
		AuthorID:          authorSlice,
	}
	result, err := repo.booksCollection.InsertOne(ctx, model)
	if err != nil {
		return &serialize.Book{}, err
	}
	if result.InsertedID != nil {
		err := repo.booksCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&model)
		if err != nil {
			return &serialize.Book{}, err
		}
	}

	categories := make([]serialize.Category, len(model.CategoryIDs))
	for i, cate := range model.CategoryIDs {
		categories[i] = serialize.Category{
			Id:      cate.Id,
			CatName: cate.CatName,
		}
	}

	authors := make([]serialize.Author, len(model.AuthorID))
	for i, author := range model.AuthorID {
		authors[i] = serialize.Author{
			Id:          author.Id,
			AuthorName:  author.AuthorName,
			DateOfBirth: author.DateOfBirth,
			HomeTown:    author.HomeTown,
			Alive:       author.Alive,
		}
	}
	return &serialize.Book{
		Id:                model.Id,
		BookName:          model.BookName,
		Price:             model.Price,
		PublishingCompany: model.PublishingCompany,
		PublicationDate:   model.PublicationDate,
		Description:       model.Description,
		CategoryIDs:       categories,
		AuthorID:          authors,
	}, nil
}

func (repo *BookRepository) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	var books []model.Book
	cursor, err := repo.booksCollection.Find(ctx, bson.M{})
	if err != nil {
		return []model.Book{}, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book model.Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (repo *BookRepository) GetBookByID(cxt context.Context, bookID string) (model.Book, error) {
	var book model.Book
	objID, _ := primitive.ObjectIDFromHex(bookID)
	err := repo.booksCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&book)
	if err != nil {
		return model.Book{}, err
	}
	return book, err
}

func (repo *BookRepository) DeleteBook(ctx context.Context, bookID string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(bookID)
	result, err := repo.booksCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return "Deleted failed", err
	}
	if result.DeletedCount == 0 {
		return "Deleted fail", err
	}
	return "Deleted successfully", nil
}

func (repo *BookRepository) EditBook(ctx context.Context, bookID string, book *serialize.Book) (*serialize.Book, error) {
	categorySlice := make([]model.Category, len(book.CategoryIDs))
	for i, cate := range book.CategoryIDs {
		categorySlice[i] = model.Category{
			Id:      cate.Id,
			CatName: cate.CatName,
		}
	}

	authorSlice := make([]model.Author, len(book.AuthorID))
	for i, author := range book.AuthorID {
		authorSlice[i] = model.Author{
			Id:          author.Id,
			AuthorName:  author.AuthorName,
			DateOfBirth: author.DateOfBirth,
			HomeTown:    author.HomeTown,
			Alive:       author.Alive,
		}
	}
	model := model.Book{
		Id:                book.Id,
		BookName:          book.BookName,
		Price:             book.Price,
		PublishingCompany: book.PublishingCompany,
		PublicationDate:   book.PublicationDate,
		Description:       book.Description,
		CategoryIDs:       categorySlice,
		AuthorID:          authorSlice,
	}
	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return &serialize.Book{}, err
	}

	result, err := repo.booksCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": model})
	if err != nil {
		return &serialize.Book{}, err
	}

	var updatedBook *serialize.Book
	if result.MatchedCount == 1 {
		err := repo.booksCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&updatedBook)
		if err != nil {
			return &serialize.Book{}, err
		}
	}
	return &serialize.Book{
		Id:                objID,
		BookName:          updatedBook.BookName,
		Price:             updatedBook.Price,
		PublishingCompany: updatedBook.PublishingCompany,
		PublicationDate:   updatedBook.PublicationDate,
		Description:       updatedBook.Description,
		CategoryIDs:       updatedBook.CategoryIDs,
		AuthorID:          updatedBook.AuthorID,
	}, nil
}
