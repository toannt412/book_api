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

var ordersCollection *mongo.Collection = configs.GetCollection(configs.DB, "orders")

func CreateOrder(cxt context.Context, order *serialize.Order) (model.Order, error) {
	result, err := ordersCollection.InsertOne(cxt, order)
	if err != nil {
		return model.Order{}, err
	}
	if result.InsertedID != nil {
		err := ordersCollection.FindOne(cxt, bson.M{"_id": result.InsertedID}).Decode(&order)
		if err != nil {
			return model.Order{}, err
		}
	}
	bookSlice := make([]model.OrderBook, len(order.Books))
	for i, book := range order.Books {
		bookSlice[i] = model.OrderBook{
			BookID:   book.BookID,
			BookName: book.BookName,
			Price:    book.Price,
			Quantity: book.Quantity,
			Total:    book.Total,
		}
	}
	return model.Order{
		Id:     order.Id,
		UserID: order.UserID,
		Books:  bookSlice,
		//CartID:        order.CartID,
		TotalQuantity: order.TotalQuantity,
		TotalPrice:    order.TotalPrice,
		TotalAmount:   order.TotalAmount,
		Status:        order.Status,
	}, nil
}

func GetOrderByID(cxt context.Context, orderID string) (model.Order, error) {
	var order model.Order
	objID, _ := primitive.ObjectIDFromHex(orderID)
	err := ordersCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&order)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func DeleteOrder(cxt context.Context, orderID string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(orderID)
	result, err := ordersCollection.DeleteOne(cxt, bson.M{"_id": objID})
	if err != nil {
		return "Deleted failed", err
	}
	if result.DeletedCount == 0 {
		return "", mongo.ErrNoDocuments
	}
	return "Deleted successfully", nil
}

func EditOrder(cxt context.Context, orderID string, order *serialize.Order) (model.Order, error) {
	objID, _ := primitive.ObjectIDFromHex(orderID)
	result, err := ordersCollection.UpdateOne(cxt, bson.M{"_id": objID}, bson.M{"$set": order})
	if err != nil {
		return model.Order{}, err
	}
	var updatedOrder model.Order
	if result.MatchedCount == 1 {
		err := ordersCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&updatedOrder)
		if err != nil {
			return model.Order{}, err
		}
	}
	return model.Order{
		Id:            updatedOrder.Id,
		UserID:        updatedOrder.UserID,
		Books:         updatedOrder.Books,
		TotalQuantity: updatedOrder.TotalQuantity,
		TotalPrice:    updatedOrder.TotalPrice,
		TotalAmount:   updatedOrder.TotalAmount,
		Status:        updatedOrder.Status,
		CartID:        updatedOrder.CartID,
	}, nil
}
