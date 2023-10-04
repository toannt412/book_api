package book

import (
	"bookstore/configs"
	"bookstore/dao/book/model"
	"bookstore/serialize"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var booksCollection *mongo.Collection = configs.GetCollection(configs.DB, "books")

func CreateBook(ctx context.Context, newBook model.Book) (model.Book, error) {
	_, err := booksCollection.InsertOne(ctx, newBook)
	if err != nil {
		return model.Book{}, err
	}
	return model.Book{
		Id:                newBook.Id,
		BookName:          newBook.BookName,
		Price:             newBook.Price,
		PublishingCompany: newBook.PublishingCompany,
		PublicationDate:   newBook.PublicationDate,
		Description:       newBook.Description,
		CategoryIDs:       newBook.CategoryIDs,
		AuthorID:          newBook.AuthorID,
	}, nil
}

func GetAllBooks(ctx context.Context) ([]model.Book, error) {
	var books []model.Book
	cursor, err := booksCollection.Find(ctx, bson.M{})
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

func GetBookByID(cxt context.Context, bookID string) (model.Book, error) {
	var book model.Book
	objID, _ := primitive.ObjectIDFromHex(bookID)
	err := booksCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&book)
	if err != nil {
		return model.Book{}, err
	}
	return book, err
}

func DeleteBook(ctx context.Context, bookID string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(bookID)
	result, err := booksCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return "Deleted failed", err
	}
	if result.DeletedCount == 0 {
		return "Deleted fail", err
	}
	return "Deleted successfully", nil
}

func EditBook(ctx context.Context, bookID string, book *serialize.Book) (model.Book, error) {
	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return model.Book{}, err
	}

	result, err := booksCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": book})
	if err != nil {
		return model.Book{}, err
	}

	var updatedBook model.Book
	if result.MatchedCount == 1 {
		err := booksCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&updatedBook)
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
