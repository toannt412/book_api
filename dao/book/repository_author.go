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

func (repo *AuthorRepository) CreateAuthor(ctx context.Context, newAuthor *serialize.Author) (*serialize.Author, error) {
	model := model.Author{
		Id:          newAuthor.Id,
		AuthorName:  newAuthor.AuthorName,
		DateOfBirth: newAuthor.DateOfBirth,
		HomeTown:    newAuthor.HomeTown,
		Alive:       newAuthor.Alive,
	}
	result, err := repo.authorsCollection.InsertOne(ctx, model)
	if err != nil {
		return &serialize.Author{}, err
	}
	if result.InsertedID != nil {
		err := repo.authorsCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&model)
		if err != nil {
			return &serialize.Author{}, err
		}
	}
	return &serialize.Author{
		Id:          model.Id,
		AuthorName:  model.AuthorName,
		DateOfBirth: model.DateOfBirth,
		HomeTown:    model.HomeTown,
		Alive:       model.Alive,
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

func (repo *AuthorRepository) EditAuthor(ctx context.Context, authorID string, author *serialize.Author) (*serialize.Author, error) {
	model := model.Author{
		Id:          author.Id,
		AuthorName:  author.AuthorName,
		DateOfBirth: author.DateOfBirth,
		HomeTown:    author.HomeTown,
		Alive:       author.Alive,
	}
	objID, err := primitive.ObjectIDFromHex(author.Id.Hex())
	if err != nil {
		return &serialize.Author{}, err
	}

	_, errUpdate := repo.authorsCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": model})
	if errUpdate != nil {
		return &serialize.Author{}, errUpdate
	}

	return &serialize.Author{
		Id:          model.Id,
		AuthorName:  model.AuthorName,
		DateOfBirth: model.DateOfBirth,
		HomeTown:    model.HomeTown,
		Alive:       model.Alive,
	}, nil
}
