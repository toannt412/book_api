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

type AuthorRepository struct {
	authorsCollection *mongo.Collection
}

func NewAuthorRepository() *AuthorRepository {
	var DB *mongo.Client = dao.ConnectDB()
	return &AuthorRepository{
		authorsCollection: dao.GetCollection(DB, "authors"),
	}
}

func (repo *AuthorRepository) CreateAuthor(ctx context.Context, newAuthor *serialize.Author) (model.Author, error) {
	result, err := repo.authorsCollection.InsertOne(ctx, newAuthor)
	if err != nil {
		return model.Author{}, err
	}
	if result.InsertedID != nil {
		err := repo.authorsCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newAuthor)
		if err != nil {
			return model.Author{}, err
		}
	}
	return model.Author{
		Id:          newAuthor.Id,
		AuthorName:  newAuthor.AuthorName,
		DateOfBirth: newAuthor.DateOfBirth,
		HomeTown:    newAuthor.HomeTown,
		Alive:       newAuthor.Alive,
	}, nil
}

func (repo *AuthorRepository) GetAuthorByID(ctx context.Context, authorID string) (model.Author, error) {
	var author model.Author
	objID, _ := primitive.ObjectIDFromHex(authorID)
	err := repo.authorsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&author)
	if err != nil {
		return model.Author{}, err
	}
	return author, err
}

func (repo *AuthorRepository) GetAllAuthors(ctx context.Context) ([]model.Author, error) {
	var authors []model.Author
	cursor, err := repo.authorsCollection.Find(ctx, bson.M{})
	if err != nil {
		return []model.Author{}, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var author model.Author
		if err := cursor.Decode(&author); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (repo *AuthorRepository) DeleteAuthor(ctx context.Context, authorID string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(authorID)
	result, err := repo.authorsCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return "Deleted fail", err
	}
	if result.DeletedCount == 0 {
		return "Deleted fail", err
	}
	return "Deleted successfully", nil
}

func (repo *AuthorRepository) EditAuthor(ctx context.Context, authorID string, author *serialize.Author) (model.Author, error) {
	objID, err := primitive.ObjectIDFromHex(authorID)
	if err != nil {
		return model.Author{}, err
	}
	// opts := options.FindOneAndUpdate().SetUpsert(true)

	// er := authorsCollection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, bson.M{"$set": author}, opts).Decode(&author)
	// if er != nil {
	// 	return model.Author{}, er
	// }
	result, err := repo.authorsCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": author})
	if err != nil {
		return model.Author{}, err
	}
	if result.MatchedCount == 0 {
		return model.Author{}, mongo.ErrNoDocuments
	}

	var updatedAuthor model.Author
	if err := repo.authorsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&updatedAuthor); err != nil {
		return model.Author{}, err
	}

	return model.Author{
		Id:          updatedAuthor.Id,
		AuthorName:  updatedAuthor.AuthorName,
		DateOfBirth: updatedAuthor.DateOfBirth,
		HomeTown:    updatedAuthor.HomeTown,
		Alive:       updatedAuthor.Alive,
	}, nil
}
