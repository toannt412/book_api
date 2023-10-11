package cart

import (
	"bookstore/dao"
	"bookstore/dao/cart/model"
	"bookstore/serialize"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepository struct {
	cartsCollection *mongo.Collection
}

func NewCartRepository() *CartRepository {
	var DB *mongo.Client = dao.ConnectDB()
	return &CartRepository{
		cartsCollection: dao.GetCollection(DB, "carts"),
	}
}

func (repo *CartRepository) CreateCart(cxt context.Context, cart *serialize.Cart) (model.Cart, error) {
	result, err := repo.cartsCollection.InsertOne(cxt, cart)
	if err != nil {
		return model.Cart{}, err
	}
	if result.InsertedID != nil {
		err := repo.cartsCollection.FindOne(cxt, bson.M{"_id": result.InsertedID}).Decode(&cart)
		if err != nil {
			return model.Cart{}, err
		}
	}
	bookSlice := make([]model.CartBook, len(cart.Books))
	for i, book := range cart.Books {
		bookSlice[i] = model.CartBook{
			BookID:   book.BookID,
			BookName: book.BookName,
			Price:    book.Price,
			Quantity: book.Quantity,
			Total:    book.Total,
		}
	}
	return model.Cart{
		Id:            cart.Id,
		UserID:        cart.UserID,
		Books:         bookSlice,
		TotalQuantity: cart.TotalQuantity,
		TotalAmount:   cart.TotalAmount,
	}, nil
}

func (repo *CartRepository) GetCart(cxt context.Context, cartID string) (model.Cart, error) {
	var cart model.Cart
	objID, _ := primitive.ObjectIDFromHex(cartID)
	err := repo.cartsCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&cart)
	if err != nil {
		return model.Cart{}, err
	}
	return cart, nil
}

func (repo *CartRepository) DeleteCart(cxt context.Context, cartID string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(cartID)
	result, err := repo.cartsCollection.DeleteOne(cxt, bson.M{"_id": objID})
	if err != nil {
		return "Deleted failed", err
	}
	if result.DeletedCount == 0 {
		return "", mongo.ErrNoDocuments
	}
	return "Deleted successfully", nil
}

func (repo *CartRepository) EditCart(cxt context.Context, cartID string, cart *serialize.Cart) (model.Cart, error) {
	objID, _ := primitive.ObjectIDFromHex(cartID)
	result, err := repo.cartsCollection.UpdateOne(cxt, bson.M{"_id": objID}, bson.M{"$set": cart})
	if err != nil {
		return model.Cart{}, err
	}
	var updatedCart model.Cart
	if result.MatchedCount == 1 {
		err := repo.cartsCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&updatedCart)
		if err != nil {
			return model.Cart{}, err
		}
	}
	return model.Cart{
		Id:            updatedCart.Id,
		UserID:        updatedCart.UserID,
		Books:         updatedCart.Books,
		TotalQuantity: updatedCart.TotalQuantity,
		TotalAmount:   updatedCart.TotalAmount,
	}, nil
}
