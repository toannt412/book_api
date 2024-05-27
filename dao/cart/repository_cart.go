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

func (repo *CartRepository) CreateCart(cxt context.Context, cart *serialize.Cart) (*serialize.Cart, error) {
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
	model := model.Cart{
		Id:            cart.Id,
		UserID:        cart.UserID,
		Books:         bookSlice,
		TotalQuantity: cart.TotalQuantity,
		TotalAmount:   cart.TotalAmount,
	}
	result, err := repo.cartsCollection.InsertOne(cxt, model)
	if err != nil {
		return &serialize.Cart{}, err
	}
	if result.InsertedID != nil {
		err := repo.cartsCollection.FindOne(cxt, bson.M{"_id": result.InsertedID}).Decode(&model)
		if err != nil {
			return &serialize.Cart{}, err
		}
	}
	books := make([]serialize.CartBook, len(model.Books))
	for i, book := range model.Books {
		books[i] = serialize.CartBook{
			BookID:   book.BookID,
			BookName: book.BookName,
			Price:    book.Price,
			Quantity: book.Quantity,
			Total:    book.Total,
		}
	}
	return &serialize.Cart{
		Id:            model.Id,
		UserID:        model.UserID,
		Books:         books,
		TotalQuantity: model.TotalQuantity,
		TotalAmount:   model.TotalAmount,
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

func (repo *CartRepository) EditCart(cxt context.Context, cartID string, cart *serialize.Cart) (*serialize.Cart, error) {
	objID, _ := primitive.ObjectIDFromHex(cartID)
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
	model := model.Cart{
		Id:            objID,
		UserID:        cart.UserID,
		Books:         bookSlice,
		TotalQuantity: cart.TotalQuantity,
		TotalAmount:   cart.TotalAmount,
	}
	result, err := repo.cartsCollection.UpdateOne(cxt, bson.M{"_id": objID}, bson.M{"$set": model})
	if err != nil {
		return &serialize.Cart{}, err
	}
	var updatedCart serialize.Cart
	if result.MatchedCount == 1 {
		err := repo.cartsCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&updatedCart)
		if err != nil {
			return &serialize.Cart{}, err
		}
	}
	books := make([]serialize.CartBook, len(model.Books))
	for i, book := range model.Books {
		books[i] = serialize.CartBook{
			BookID:   book.BookID,
			BookName: book.BookName,
			Price:    book.Price,
			Quantity: book.Quantity,
			Total:    book.Total,
		}
	}
	return &serialize.Cart{
		Id:            model.Id,
		UserID:        model.UserID,
		Books:         books,
		TotalQuantity: model.TotalQuantity,
		TotalAmount:   model.TotalAmount,
	}, nil
}
