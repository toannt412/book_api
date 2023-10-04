package cart

import (
	"bookstore/configs"
	"bookstore/dao/cart/model"
	"bookstore/serialize"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartsCollection *mongo.Collection = configs.GetCollection(configs.DB, "carts")

func CreateCart(cxt context.Context, cart *serialize.Cart) (model.Cart, error) {
	result, err := cartsCollection.InsertOne(cxt, cart)
	if err != nil {
		return model.Cart{}, err
	}
	if result.InsertedID != nil {
		err := cartsCollection.FindOne(cxt, bson.M{"_id": result.InsertedID}).Decode(&cart)
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

func GetCart(cxt context.Context, cartID string) (model.Cart, error) {
	var cart model.Cart
	objID, _ := primitive.ObjectIDFromHex(cartID)
	err := cartsCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&cart)
	if err != nil {
		return model.Cart{}, err
	}
	return cart, nil
}

func DeleteCart(cxt context.Context, cartID string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(cartID)
	result, err := cartsCollection.DeleteOne(cxt, bson.M{"_id": objID})
	if err != nil {
		return "Deleted failed", err
	}
	if result.DeletedCount == 0 {
		return "", mongo.ErrNoDocuments
	}
	return "Deleted successfully", nil
}

func EditCart(cxt context.Context, cartID string, cart *serialize.Cart) (model.Cart, error) {
	objID, _ := primitive.ObjectIDFromHex(cartID)
	result, err := cartsCollection.UpdateOne(cxt, bson.M{"_id": objID}, bson.M{"$set": cart})
	if err != nil {
		return model.Cart{}, err
	}
	var updatedCart model.Cart
	if result.MatchedCount == 1 {
		err := cartsCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&updatedCart)
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
