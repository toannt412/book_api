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

func (repo *BookRepository) CreateBook(ctx context.Context, newBook *serialize.Book) (model.Book, error) {
	result, err := repo.booksCollection.InsertOne(ctx, newBook)
	if err != nil {
		return model.Book{}, err
	}
	if result.InsertedID != nil {
		err := repo.booksCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newBook)
		if err != nil {
			return model.Book{}, err
		}
	}

	categories := make([]model.Category, len(newBook.CategoryIDs))
	for i, cate := range newBook.CategoryIDs {
		categories[i] = model.Category{
			Id:      cate.Id,
			CatName: cate.CatName,
		}
	}

	authors := make([]model.Author, len(newBook.AuthorID))
	for i, author := range newBook.AuthorID {
		authors[i] = model.Author{
			Id:          author.Id,
			AuthorName:  author.AuthorName,
			DateOfBirth: author.DateOfBirth,
			HomeTown:    author.HomeTown,
			Alive:       author.Alive,
		}
	}
	return model.Book{
		Id:                newBook.Id,
		BookName:          newBook.BookName,
		Price:             newBook.Price,
		PublishingCompany: newBook.PublishingCompany,
		PublicationDate:   newBook.PublicationDate,
		Description:       newBook.Description,
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

func (repo *BookRepository) EditBook(ctx context.Context, bookID string, book *serialize.Book) (model.Book, error) {
	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return model.Book{}, err
	}

	result, err := repo.booksCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": book})
	if err != nil {
		return model.Book{}, err
	}

	var updatedBook model.Book
	if result.MatchedCount == 1 {
		err := repo.booksCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&updatedBook)
		if err != nil {
			return model.Book{}, err
		}
	}
	return model.Book{
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
