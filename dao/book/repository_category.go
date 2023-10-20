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

type CategoryRepository struct {
	categoriesCollection *mongo.Collection
}

func NewCategoryRepository() *CategoryRepository {
	var DB *mongo.Client = dao.ConnectDB()
	return &CategoryRepository{
		categoriesCollection: dao.GetCollection(DB, "categories"),
	}
}

func (repo *CategoryRepository) CreateCategory(ctx context.Context, newCategory *serialize.Category) (*serialize.Category, error) {
	model := model.Category{
		Id:      newCategory.Id,
		CatName: newCategory.CatName,
	}
	result, err := repo.categoriesCollection.InsertOne(ctx, model)
	if err != nil {
		return &serialize.Category{}, err
	}
	if result.InsertedID != nil {
		err := repo.categoriesCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&model)
		if err != nil {
			return &serialize.Category{}, err
		}
	}
	return &serialize.Category{
		Id:      model.Id,
		CatName: model.CatName,
	}, nil
}

func (repo *CategoryRepository) GetCategoryByID(ctx context.Context, categoryID string) (model.Category, error) {
	var category model.Category
	objID, _ := primitive.ObjectIDFromHex(categoryID)
	err := repo.categoriesCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&category)
	if err != nil {
		return model.Category{}, err
	}
	return category, err
}

func (repo *CategoryRepository) GetAllCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	cursor, err := repo.categoriesCollection.Find(ctx, bson.M{})
	if err != nil {
		return []model.Category{}, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var category model.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (repo *CategoryRepository) DeleteCategory(ctx context.Context, categoryID string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(categoryID)
	result, err := repo.categoriesCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return "Deleted failed", err
	}
	if result.DeletedCount == 0 {
		return "Deleted fail", nil
	}
	return "Deleted successfully", err
}

func (repo *CategoryRepository) EditCategory(ctx context.Context, categoryID string, category *serialize.Category) (*serialize.Category, error) {
	objID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return &serialize.Category{}, err
	}
	update := bson.M{"_id": objID, "categoryname": category.CatName}
	result, err := repo.categoriesCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
	if err != nil {
		return &serialize.Category{}, err
	}
	if result.ModifiedCount == 0 {
		return &serialize.Category{}, err
	}
	return &serialize.Category{
		Id:      objID,
		CatName: category.CatName,
	}, nil
}
