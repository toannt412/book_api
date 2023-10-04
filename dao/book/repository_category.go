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

var categoriesCollection *mongo.Collection = configs.GetCollection(configs.DB, "categories")

func CreateCategory(ctx context.Context, newCategory *serialize.Category) (model.Category, error) {
	result, err := categoriesCollection.InsertOne(ctx, newCategory)
	if err != nil {
		return model.Category{}, err
	}
	if result.InsertedID != nil {
		err := categoriesCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newCategory)
		if err != nil {
			return model.Category{}, err
		}
	}
	return model.Category{
		Id:      newCategory.Id,
		CatName: newCategory.CatName,
	}, nil
}

func GetCategoryByID(ctx context.Context, categoryID string) (model.Category, error) {
	var category model.Category
	objID, _ := primitive.ObjectIDFromHex(categoryID)
	err := categoriesCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&category)
	if err != nil {
		return model.Category{}, err
	}
	return category, err
}

func GetAllCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	cursor, err := categoriesCollection.Find(ctx, bson.M{})
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

func DeleteCategory(ctx context.Context, categoryID string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(categoryID)
	result, err := categoriesCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return "Deleted failed", err
	}
	if result.DeletedCount == 0 {
		return "Deleted fail", nil
	}
	return "Deleted successfully", err
}

func EditCategory(ctx context.Context, categoryID string, category *serialize.Category) (model.Category, error) {
	objID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return model.Category{}, err
	}
	update := bson.M{"_id": objID, "categoryname": category.CatName}
	result, err := categoriesCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
	if err != nil {
		return model.Category{}, err
	}
	if result.ModifiedCount == 0 {
		return model.Category{}, err
	}
	return model.Category{
		Id:      objID,
		CatName: category.CatName,
	}, nil
}

